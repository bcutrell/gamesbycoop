package screens

import (
	"math/rand"

	"jumponblocks/sprites"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	gravity        = 0.5
	jumpVelocity   = -12.0
	moveSpeed      = 5.0
	blockWidth     = 80
	blockHeight    = 20
	spikeThickness = 15
)

type Player struct {
	X, Y     float64
	VelX     float64
	VelY     float64
	Width    float64
	Height   float64
	OnGround bool
}

type Block struct {
	X, Y       float64
	Width      float64
	Height     float64
	VelY       float64
	IsStarting bool // True if this is the starting block
}

type GameplayScreen struct {
	player              *Player
	blocks              []*Block
	blockImage          *ebiten.Image
	character           *ebiten.Image
	characterType       sprites.CharacterType
	score               float64
	spawnTimer          int
	blockSpeed          float64
	gameOver            bool
	onGameOver          func(score int)
	startingBlockActive bool // Whether starting block is still present
	wasOnStartingBlock  bool // Track if player was on starting block
}

func NewGameplayScreen(onGameOver func(score int)) *GameplayScreen {
	gs := &GameplayScreen{
		blocks:        make([]*Block, 0),
		blockImage:    sprites.NewBlock(blockWidth, blockHeight),
		character:     sprites.NewSmileyFace(),
		characterType: sprites.CharacterSmiley,
		blockSpeed:    1.5,
		onGameOver:    onGameOver,
	}

	gs.spawnInitialBlocks()

	return gs
}

func (gs *GameplayScreen) SetCharacter(charType sprites.CharacterType) {
	gs.characterType = charType
	gs.character = sprites.NewCharacter(charType)
}

func (gs *GameplayScreen) spawnInitialBlocks() {
	// Create starting platform in the middle-bottom area
	startBlockY := float64(ScreenHeight) - 150
	startBlockX := float64(ScreenWidth)/2 - blockWidth/2

	gs.blocks = append(gs.blocks, &Block{
		X:          startBlockX,
		Y:          startBlockY,
		Width:      blockWidth,
		Height:     blockHeight,
		VelY:       0, // Starting block doesn't move initially
		IsStarting: true,
	})

	// Place player ON the starting block
	gs.player = &Player{
		X:        startBlockX + blockWidth/2 - 30,
		Y:        startBlockY - 60, // On top of block
		Width:    60,
		Height:   60,
		OnGround: true,
	}

	gs.startingBlockActive = true
	gs.wasOnStartingBlock = true

	// Spawn additional blocks at various heights for jumping
	blockPositions := []struct{ x, y float64 }{
		{float64(spikeThickness) + 30, startBlockY - 120},
		{float64(ScreenWidth) - spikeThickness - blockWidth - 30, startBlockY - 80},
		{float64(ScreenWidth)/2 - blockWidth/2, startBlockY - 200},
		{float64(spikeThickness) + 80, startBlockY - 280},
		{float64(ScreenWidth) - spikeThickness - blockWidth - 80, startBlockY - 350},
	}

	for _, pos := range blockPositions {
		gs.blocks = append(gs.blocks, &Block{
			X:      pos.x,
			Y:      pos.y,
			Width:  blockWidth,
			Height: blockHeight,
			VelY:   gs.blockSpeed,
		})
	}
}

func (gs *GameplayScreen) Reset() {
	gs.blocks = make([]*Block, 0)
	gs.score = 0
	gs.spawnTimer = 0
	gs.blockSpeed = 1.5
	gs.gameOver = false
	gs.startingBlockActive = false
	gs.wasOnStartingBlock = false
	// Refresh character in case it changed
	gs.character = sprites.NewCharacter(gs.characterType)

	gs.spawnInitialBlocks()
}

func (gs *GameplayScreen) removeStartingBlock() {
	if !gs.startingBlockActive {
		return
	}
	newBlocks := make([]*Block, 0, len(gs.blocks))
	for _, b := range gs.blocks {
		if !b.IsStarting {
			newBlocks = append(newBlocks, b)
		}
	}
	gs.blocks = newBlocks
	gs.startingBlockActive = false
}

func (gs *GameplayScreen) spawnBlock() {
	playAreaWidth := float64(ScreenWidth) - 2*spikeThickness - blockWidth
	x := float64(spikeThickness) + rand.Float64()*playAreaWidth

	gs.blocks = append(gs.blocks, &Block{
		X:      x,
		Y:      -blockHeight,
		Width:  blockWidth,
		Height: blockHeight,
		VelY:   gs.blockSpeed,
	})
}

func (gs *GameplayScreen) Update() {
	if gs.gameOver {
		return
	}

	// Handle input
	gs.player.VelX = 0
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		gs.player.VelX = -moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		gs.player.VelX = moveSpeed
	}
	if (ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW)) && gs.player.OnGround {
		gs.player.VelY = jumpVelocity
		gs.player.OnGround = false
	}

	// Apply gravity
	gs.player.VelY += gravity

	// Move player
	gs.player.X += gs.player.VelX
	gs.player.Y += gs.player.VelY

	// Update blocks
	gs.spawnTimer++
	if gs.spawnTimer >= 60 { // Spawn every 1 second at 60 TPS
		gs.spawnBlock()
		gs.spawnTimer = 0
	}

	// Move blocks and check for removal
	newBlocks := make([]*Block, 0, len(gs.blocks))
	for _, b := range gs.blocks {
		b.Y += b.VelY
		// Keep blocks that are still on screen
		if b.Y < float64(ScreenHeight)+blockHeight {
			newBlocks = append(newBlocks, b)
		}
	}
	gs.blocks = newBlocks

	// Check block collisions (landing on blocks)
	gs.player.OnGround = false
	onStartingBlock := false
	for _, b := range gs.blocks {
		if gs.checkBlockCollision(b) {
			gs.player.OnGround = true
			gs.player.VelY = b.VelY // Move with the block
			gs.player.Y = b.Y - gs.player.Height
			if b.IsStarting {
				onStartingBlock = true
			}
		}
	}

	// Check if player jumped off starting block OR score reached 5
	if gs.startingBlockActive {
		// Player was on starting block but isn't anymore (jumped off)
		if gs.wasOnStartingBlock && !onStartingBlock {
			gs.removeStartingBlock()
		}
		// Score reached 5
		if gs.score >= 5.0 {
			gs.removeStartingBlock()
		}
	}
	gs.wasOnStartingBlock = onStartingBlock

	// Check spike collisions (edges)
	if gs.checkSpikeCollision() {
		gs.gameOver = true
		if gs.onGameOver != nil {
			gs.onGameOver(int(gs.score))
		}
		return
	}

	// Update score (time survived)
	gs.score += 1.0 / 60.0 // 1 point per second at 60 TPS

	// Gradually increase difficulty based on score
	gs.blockSpeed = min(1.5+gs.score*0.05, 6.0)
}

func (gs *GameplayScreen) checkBlockCollision(b *Block) bool {
	// Check if player is falling onto the block from above
	playerBottom := gs.player.Y + gs.player.Height
	playerRight := gs.player.X + gs.player.Width
	blockRight := b.X + b.Width

	// Horizontal overlap
	if gs.player.X >= blockRight || playerRight <= b.X {
		return false
	}

	// Check if landing on top of block (falling down and near the top)
	if gs.player.VelY >= 0 && playerBottom >= b.Y && playerBottom <= b.Y+b.Height+gs.player.VelY {
		return true
	}

	return false
}

func (gs *GameplayScreen) checkSpikeCollision() bool {
	// Check all four edges
	playLeft := float64(spikeThickness)
	playRight := float64(ScreenWidth - spikeThickness)
	playTop := float64(spikeThickness)
	playBottom := float64(ScreenHeight - spikeThickness)

	// Player bounds
	playerRight := gs.player.X + gs.player.Width
	playerBottom := gs.player.Y + gs.player.Height

	// Left edge
	if gs.player.X < playLeft {
		return true
	}
	// Right edge
	if playerRight > playRight {
		return true
	}
	// Top edge
	if gs.player.Y < playTop {
		return true
	}
	// Bottom edge
	if playerBottom > playBottom {
		return true
	}

	return false
}

func (gs *GameplayScreen) Draw(screen *ebiten.Image) {
	// Draw blocks
	for _, b := range gs.blocks {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(b.X, b.Y)
		screen.DrawImage(gs.blockImage, op)
	}

	// Draw player (character)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(gs.player.X, gs.player.Y)
	screen.DrawImage(gs.character, op)

	// Draw spikes on top
	sprites.DrawSpikes(screen, ScreenWidth, ScreenHeight, spikeThickness)
}

func (gs *GameplayScreen) GetScore() int {
	return int(gs.score)
}
