package sprites

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	spriteSize  = 15
	scaleFactor = 4
	outputSize  = spriteSize * scaleFactor
)

// Character types
type CharacterType int

const (
	CharacterSmiley CharacterType = iota
	CharacterCool
	CharacterWink
	CharacterAngry
	CharacterSurprised
	CharacterRobot
	CharacterCount // Used to count total characters
)

// CharacterNames for display in the selection screen.
var CharacterNames = []string{"Smiley", "Cool", "Wink", "Angry", "Shocked", "Robot"}

// Character patterns use ASCII art where each character maps to a color:
//   '.' = transparent (background shows through)
//   'X' = black (outline)
//   '#' = yellow (skin tone)
//   'O' = black (eyes, mouth features)
//   '%' = silver (robot metal)

var smileyPattern = []string{
	".....XXXXX.....",
	"...XX#####XX...",
	"..X#########X..",
	".X###########X.",
	"X#OOOOOOOOOOO#X",
	"X#OOOO###OOOO#X",
	"X#############X",
	"X#############X",
	"X#############X",
	"X##########O##X",
	".X####O##OO##X.",
	"..X####OO###X..",
	"...XX#####XX...",
	".....XXXXX.....",
	"...............",
}

// Cool character with sunglasses
var coolPattern = []string{
	".....XXXXX.....",
	"...XX#####XX...",
	"..X#########X..",
	".X###########X.",
	"XXXXXXXXXXXXXXX",
	"X#XXX###XXX##X.",
	"X#############X",
	"X#############X",
	"X#############X",
	"X##########O##X",
	".X####O##OO##X.",
	"..X####OO###X..",
	"...XX#####XX...",
	".....XXXXX.....",
	"...............",
}

// Wink character
var winkPattern = []string{
	".....XXXXX.....",
	"...XX#####XX...",
	"..X#########X..",
	".X###########X.",
	"X#OOO###OOOOO#X",
	"X#OOO###OOOOO#X",
	"X#############X",
	"X#############X",
	"X#############X",
	"X##########O##X",
	".X####O##OO##X.",
	"..X####OO###X..",
	"...XX#####XX...",
	".....XXXXX.....",
	"...............",
}

// Angry character (frown + angry eyebrows)
var angryPattern = []string{
	".....XXXXX.....",
	"...XX#####XX...",
	"..X#########X..",
	".X#O#######O#X.",
	"X##O#####O###X.",
	"X#OOOO#OOOOO#X.",
	"X#############X",
	"X#############X",
	"X#############X",
	"X####OOO#####X.",
	".X##O###O###X..",
	"..X#########X..",
	"...XX#####XX...",
	".....XXXXX.....",
	"...............",
}

// Surprised character (O mouth, big eyes)
var surprisedPattern = []string{
	".....XXXXX.....",
	"...XX#####XX...",
	"..X#########X..",
	".X###########X.",
	"X#OOO###OOO##X.",
	"X#OOO###OOO##X.",
	"X#############X",
	"X#############X",
	"X#####OOO####X.",
	"X####O###O###X.",
	".X###O###O##X..",
	"..X###OOO##X...",
	"...XX#####XX...",
	".....XXXXX.....",
	"...............",
}

// Robot character (square eyes, antenna)
var robotPattern = []string{
	"......XXX......",
	"......X#X......",
	".....XXXXX.....",
	"...XXXXXXXXX...",
	"..X%%%%%%%%%X..",
	".X%%%%%%%%%%%X.",
	"X%OOOO%OOOO%%X.",
	"X%OOOO%OOOO%%X.",
	"X%%%%%%%%%%%%%X",
	"X%%%%%%%%%%%X..",
	".X%%OOOOO%%X...",
	"..X%%%%%%%X....",
	"...XXXXXXX.....",
	"...............",
	"...............",
}

var charColors = map[rune]color.RGBA{
	'.': {0, 0, 0, 0},         // Transparent
	'X': {0, 0, 0, 255},       // Black outline
	'#': {255, 220, 0, 255},   // Yellow skin
	'O': {0, 0, 0, 255},       // Black features
	'%': {180, 180, 200, 255}, // Silver metal
}

func createCharacterFromPattern(pattern []string) *ebiten.Image {
	smallImg := ebiten.NewImage(spriteSize, spriteSize)

	for y, row := range pattern {
		for x, char := range row {
			if c, ok := charColors[char]; ok {
				smallImg.Set(x, y, c)
			}
		}
	}

	scaledImg := ebiten.NewImage(outputSize, outputSize)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(scaleFactor), float64(scaleFactor))
	op.Filter = ebiten.FilterNearest
	scaledImg.DrawImage(smallImg, op)

	return scaledImg
}

func NewSmileyFace() *ebiten.Image {
	return createCharacterFromPattern(smileyPattern)
}

func NewCharacter(charType CharacterType) *ebiten.Image {
	switch charType {
	case CharacterCool:
		return createCharacterFromPattern(coolPattern)
	case CharacterWink:
		return createCharacterFromPattern(winkPattern)
	case CharacterAngry:
		return createCharacterFromPattern(angryPattern)
	case CharacterSurprised:
		return createCharacterFromPattern(surprisedPattern)
	case CharacterRobot:
		return createCharacterFromPattern(robotPattern)
	default:
		return createCharacterFromPattern(smileyPattern)
	}
}

func GetAllCharacters() []*ebiten.Image {
	chars := make([]*ebiten.Image, CharacterCount)
	for i := CharacterType(0); i < CharacterCount; i++ {
		chars[i] = NewCharacter(i)
	}
	return chars
}
