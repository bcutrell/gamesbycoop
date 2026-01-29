package screens

import (
	"fmt"
	"image/color"

	"jumponblocks/sprites"
	"jumponblocks/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	ColorGrey     = color.RGBA{128, 128, 128, 255}
	ColorGold     = color.RGBA{255, 215, 0, 255}
	ColorSelected = color.RGBA{100, 200, 100, 255}
	ColorOwned    = color.RGBA{200, 200, 200, 255}
)

type CharacterInfo struct {
	CharType sprites.CharacterType
	Price    float64
	Owned    bool
	Image    *ebiten.Image
}

type CharSelectScreen struct {
	characters   []CharacterInfo
	selectedChar sprites.CharacterType
	backButton   *ui.Button
	onBack       func()
	onSelect     func(sprites.CharacterType)
	wallet       *float64 // Pointer to shared wallet
	clickLock    bool
}

func NewCharSelectScreen(wallet *float64, onBack func(), onSelect func(sprites.CharacterType)) *CharSelectScreen {
	cs := &CharSelectScreen{
		characters:   make([]CharacterInfo, int(sprites.CharacterCount)),
		selectedChar: sprites.CharacterSmiley,
		wallet:       wallet,
		onBack:       onBack,
		onSelect:     onSelect,
	}

	// Initialize characters with prices
	// Smiley is free, others cost money
	prices := []float64{0.00, 1.00, 2.50, 3.00, 4.50, 10.00}

	for i := sprites.CharacterType(0); i < sprites.CharacterCount; i++ {
		cs.characters[i] = CharacterInfo{
			CharType: i,
			Price:    prices[i],
			Owned:    prices[i] == 0, // Free characters start owned
			Image:    sprites.NewCharacter(i),
		}
	}

	// Back button
	cs.backButton = ui.NewButton(
		"Back",
		20, 20,
		80, 40,
		ColorButton,
		ColorButtonHover,
		ColorText,
	)
	cs.backButton.OnClick = func() {
		if !cs.clickLock {
			cs.clickLock = true
			if onBack != nil {
				onBack()
			}
		}
	}

	return cs
}

func (cs *CharSelectScreen) SetSelectedChar(charType sprites.CharacterType) {
	cs.selectedChar = charType
}

func (cs *CharSelectScreen) Update() {
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		cs.clickLock = false
	}

	cs.backButton.Update()

	// Check for character clicks
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !cs.clickLock {
		mx, my := ebiten.CursorPosition()
		cs.handleCharacterClick(mx, my)
	}
}

func (cs *CharSelectScreen) handleCharacterClick(mx, my int) {
	// Grid layout: 3 characters in a row
	startX := 50
	startY := 120
	cellWidth := 110
	cellHeight := 140

	for i := range cs.characters {
		col := i % 3
		row := i / 3

		x := startX + col*cellWidth
		y := startY + row*cellHeight

		// Check if click is within this character's cell
		if mx >= x && mx < x+cellWidth-10 && my >= y && my < y+cellHeight-10 {
			cs.clickLock = true
			char := &cs.characters[i]

			if char.Owned {
				// Select this character
				cs.selectedChar = char.CharType
				if cs.onSelect != nil {
					cs.onSelect(char.CharType)
				}
			} else if *cs.wallet >= char.Price {
				// Buy this character
				*cs.wallet -= char.Price
				char.Owned = true
				cs.selectedChar = char.CharType
				if cs.onSelect != nil {
					cs.onSelect(char.CharType)
				}
			}
			break
		}
	}
}

func (cs *CharSelectScreen) Draw(screen *ebiten.Image, face text.Face) {
	// Title
	titleOp := &text.DrawOptions{}
	title := "Select Character"
	tw, _ := text.Measure(title, face, 0)
	titleOp.GeoM.Translate(float64(ScreenWidth)/2-tw/2, 70)
	titleOp.ColorScale.ScaleWithColor(ColorText)
	text.Draw(screen, title, face, titleOp)

	// Wallet display
	walletMsg := fmt.Sprintf("Wallet: $%.2f", *cs.wallet)
	walletOp := &text.DrawOptions{}
	ww, _ := text.Measure(walletMsg, face, 0)
	walletOp.GeoM.Translate(float64(ScreenWidth)-ww-20, 35)
	walletOp.ColorScale.ScaleWithColor(ColorGold)
	text.Draw(screen, walletMsg, face, walletOp)

	// Draw character grid
	startX := 50
	startY := 120
	cellWidth := 110
	cellHeight := 140

	for i, char := range cs.characters {
		col := i % 3
		row := i / 3

		x := float64(startX + col*cellWidth)
		y := float64(startY + row*cellHeight)

		// Draw cell background
		bgColor := color.RGBA{60, 60, 80, 255}
		if char.CharType == cs.selectedChar {
			bgColor = ColorSelected
		} else if char.Owned {
			bgColor = color.RGBA{80, 80, 100, 255}
		}
		vector.FillRect(screen, float32(x), float32(y), float32(cellWidth-10), float32(cellHeight-10), bgColor, false)

		// Draw character border
		borderColor := color.RGBA{100, 100, 120, 255}
		if char.CharType == cs.selectedChar {
			borderColor = ColorGold
		}
		vector.StrokeRect(screen, float32(x), float32(y), float32(cellWidth-10), float32(cellHeight-10), 2, borderColor, false)

		// Draw character image (scaled down)
		charW := char.Image.Bounds().Dx()
		op := &ebiten.DrawImageOptions{}
		scale := 0.7
		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(
			x+float64(cellWidth-10)/2-float64(charW)*scale/2,
			y+15,
		)
		screen.DrawImage(char.Image, op)

		// Draw character name
		nameOp := &text.DrawOptions{}
		name := sprites.CharacterNames[i]
		nw, _ := text.Measure(name, face, 0)
		nameOp.GeoM.Translate(x+float64(cellWidth-10)/2-nw/2, y+75)
		nameOp.ColorScale.ScaleWithColor(ColorText)
		text.Draw(screen, name, face, nameOp)

		// Draw price or "Owned" status
		var priceText string
		var priceColor color.Color

		if char.Owned {
			if char.CharType == cs.selectedChar {
				priceText = "Selected"
				priceColor = ColorGold
			} else {
				priceText = "Owned"
				priceColor = ColorOwned
			}
		} else {
			priceText = fmt.Sprintf("$%.2f", char.Price)
			if *cs.wallet >= char.Price {
				priceColor = ColorGold
			} else {
				priceColor = ColorGrey
			}
		}

		priceOp := &text.DrawOptions{}
		pw, _ := text.Measure(priceText, face, 0)
		priceOp.GeoM.Translate(x+float64(cellWidth-10)/2-pw/2, y+95)
		priceOp.ColorScale.ScaleWithColor(priceColor)
		text.Draw(screen, priceText, face, priceOp)

		// Draw "Buy" hint for affordable unowned characters
		if !char.Owned && *cs.wallet >= char.Price {
			buyOp := &text.DrawOptions{}
			buyText := "Click to Buy"
			bw, _ := text.Measure(buyText, face, 0)
			buyOp.GeoM.Translate(x+float64(cellWidth-10)/2-bw/2, y+115)
			buyOp.ColorScale.ScaleWithColor(ColorGold)
			text.Draw(screen, buyText, face, buyOp)
		}
	}

	// Back button
	cs.backButton.Draw(screen, face)
}
