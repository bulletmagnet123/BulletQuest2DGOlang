package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Additional GameCharacter constants (Player is defined in Player.go)
const (
	GameCharacterSkeleton GameCharacter = 1
)

const (
	SPRITE_DEFAULT_SIZE = 16
	SPRITE_ROWS         = 7
	SPRITE_COLS         = 4
)

type CharacterSprites struct {
	SpriteSheet *ebiten.Image
	Sprites     [][]*ebiten.Image
}

var GameCharactersData = make(map[GameCharacter]*CharacterSprites)

// LoadGameCharacters reads sprite sheets from disk and slices them into frames.
// It expects files at the paths provided in the map below. Change paths if needed.
func LoadGameCharacters() {
	files := map[GameCharacter]string{
		GameCharacterPlayer:   "assets/playersheet.png",
		GameCharacterSkeleton: "assets/skeletonsheet.png",
	}

	for gc, path := range files {
		sheet, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			// If file missing, log and continue with nil entry
			log.Printf("warning: could not load %s: %v", path, err)
			GameCharactersData[gc] = &CharacterSprites{}
			continue
		}

		sprites := make([][]*ebiten.Image, SPRITE_ROWS)
		for r := 0; r < SPRITE_ROWS; r++ {
			sprites[r] = make([]*ebiten.Image, SPRITE_COLS)
			for c := 0; c < SPRITE_COLS; c++ {
				x0 := c * SPRITE_DEFAULT_SIZE
				y0 := r * SPRITE_DEFAULT_SIZE
				rect := image.Rect(x0, y0, x0+SPRITE_DEFAULT_SIZE, y0+SPRITE_DEFAULT_SIZE)
				sub := sheet.SubImage(rect).(*ebiten.Image)
				sprites[r][c] = sub
			}
		}

		GameCharactersData[gc] = &CharacterSprites{
			SpriteSheet: sheet,
			Sprites:     sprites,
		}
	}
}

func (gc GameCharacter) GetAnimationFrames(direction int) int {
	// Default: 4 frames for current characters
	switch gc {
	case GameCharacterPlayer:
		return 4
	case GameCharacterSkeleton:
		return 4
	default:
		return 1
	}
}

func (gc GameCharacter) GetSpriteSheet() *ebiten.Image {
	if data, ok := GameCharactersData[gc]; ok {
		return data.SpriteSheet
	}
	return nil
}

func (gc GameCharacter) GetSprite(yPos, xPos int) *ebiten.Image {
	if data, ok := GameCharactersData[gc]; ok {
		if yPos >= 0 && yPos < len(data.Sprites) {
			row := data.Sprites[yPos]
			if xPos >= 0 && xPos < len(row) {
				return row[xPos]
			}
		}
	}
	return nil
}
