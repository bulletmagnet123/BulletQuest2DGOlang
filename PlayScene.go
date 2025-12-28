package main

import (
	"image"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// PlayScene: simple gameplay placeholder moved to its own file
type PlayScene struct {
	sm         *SceneManager
	Player     *Player
	MapManager interface {
		Draw(screen *ebiten.Image)
		CanMoveHere(x, y float64) bool
	}
	Skeleton    *Character
	PlayingUI   interface{ DrawUI(screen *ebiten.Image) }
	BtnExit     *CustomButton
	tilemapJSON *TilemapJSON
	tilemapImg  *ebiten.Image
}

func NewPlayScene(sm *SceneManager) *PlayScene {
	p := &PlayScene{sm: sm, Player: NewPlayer()}

	// attempt to load the tilemap JSON and tileset image for the PlayScene
	if tm, err := NewTilemapJSON("assets/maps/dirtmap.json"); err != nil {
		log.Println("failed to load tilemap JSON:", err)
	} else {
		p.tilemapJSON = tm
	}

	if img, _, err := ebitenutil.NewImageFromFile("assets/maps/floorsheet.png"); err != nil {
		log.Println("failed to load tilemap image:", err)
	} else {
		p.tilemapImg = img
	}

	return p

}

func (p *PlayScene) Enter() {}
func (p *PlayScene) Exit()  {}

func (p *PlayScene) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		p.sm.GoTo(NewMenuScene(p.sm))
	}
	// simple fixed delta (approx 60 FPS). Replace with real delta if available.
	delta := 1.0 / 60.0
	p.updatePlayerMove(delta)
	return nil

}

// updatePlayerMove moves the player using WASD and sets facing direction; F sets attacking flag.
func (p *PlayScene) updatePlayerMove(delta float64) {
	if p.Player == nil {
		return
	}

	// read input
	up := ebiten.IsKeyPressed(ebiten.KeyW)
	down := ebiten.IsKeyPressed(ebiten.KeyS)
	left := ebiten.IsKeyPressed(ebiten.KeyA)
	right := ebiten.IsKeyPressed(ebiten.KeyD)
	attack := ebiten.IsKeyPressed(ebiten.KeyF)

	p.Player.Attacking = attack

	dx := 0.0
	dy := 0.0
	if right && !left {
		dx = 1
	} else if left && !right {
		dx = -1
	}
	if down && !up {
		dy = 1
	} else if up && !down {
		dy = -1
	}

	// if no movement keys pressed, reset animation and return
	if dx == 0 && dy == 0 {
		p.Player.ResetAnimation()
		return
	}

	// compute normalized movement similar to original algorithm
	// baseSpeed = delta * 300
	baseSpeed := delta * 150

	// avoid divide by zero
	ratio := 0.0
	if dx != 0 {
		ratio = math.Abs(dy) / math.Abs(dx)
	} else {
		ratio = 1e6
	}
	angle := math.Atan(ratio)
	xSpeed := math.Cos(angle)
	ySpeed := math.Sin(angle)

	// determine facing based on larger component
	if xSpeed > ySpeed {
		if dx > 0 {
			p.Player.SetFaceDir(FACE_DIR_RIGHT)
		} else {
			p.Player.SetFaceDir(FACE_DIR_LEFT)
		}
	} else {
		if dy > 0 {
			p.Player.SetFaceDir(FACE_DIR_DOWN)
		} else {
			p.Player.SetFaceDir(FACE_DIR_UP)
		}
	}

	if dx < 0 {
		xSpeed *= -1
	}
	if dy < 0 {
		ySpeed *= -1
	}

	deltaX := xSpeed * baseSpeed
	deltaY := ySpeed * baseSpeed

	// proposed new position
	newX := p.Player.Position.X + deltaX
	newY := p.Player.Position.Y + deltaY

	// ask map if movement allowed. If no MapManager provided, allow movement.
	canMove := true
	if p.MapManager != nil {
		canMove = p.MapManager.CanMoveHere(newX, newY)
	}

	if canMove {
		p.Player.Position.X = newX
		p.Player.Position.Y = newY
		p.Player.UpdateAnimation()
	} else {
		p.Player.ResetAnimation()
	}
}

func (p *PlayScene) Draw(screen *ebiten.Image) {
	screen.Clear()
	// Render order similar to original Java: map, player, other chars, UI, buttons

	// Draw exit button if provided
	if p.BtnExit != nil {
		p.BtnExit.Draw(screen)
	}

	if p.tilemapJSON == nil || p.tilemapImg == nil {
		// nothing to draw
		log.Println("tilemapJSON or tilemapImg is nil")
		return
	}

	// determine tiles-per-row from the spritesheet size rather than hardcoding 22
	imgW, _ := p.tilemapImg.Size()
	tilesPerRow := imgW / 16

	// reuse options to avoid allocating per-tile
	opts := &ebiten.DrawImageOptions{}

	for _, layer := range p.tilemapJSON.Layers {
		// loop over the tiles in the layer data
		for index, id := range layer.Data {
			// skip empty tiles (commonly 0 in Tiled exports)
			if id <= 0 {
				continue
			}

			// get the tile position of the tile (in tiles)
			tx := index % layer.Width
			ty := index / layer.Width

			// convert the tile position to pixel position
			px := tx * 16
			py := ty * 16

			// compute the source tile position in the spritesheet
			sCol := (id - 1) % tilesPerRow
			sRow := (id - 1) / tilesPerRow
			srcX := sCol * 16
			srcY := sRow * 16

			// set the draw options and draw
			opts.GeoM.Reset()
			opts.GeoM.Translate(float64(px), float64(py))
			screen.DrawImage(p.tilemapImg.SubImage(image.Rect(srcX, srcY, srcX+16, srcY+16)).(*ebiten.Image), opts)
		}
	}

	p.drawPlayer(screen)

	// Centered help text
	DrawTextAtCenter(screen, "Gameplay - press ESC to return")
}

func (p *PlayScene) drawPlayer(screen *ebiten.Image) {
	if p.Player == nil {
		return
	}
	gc := p.Player.GetGameCharType()
	// Java used getSprite(aniIndex, faceDir)
	sprite := gc.GetSprite(p.Player.GetAniIndex(), p.Player.GetFaceDir())
	if sprite == nil {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Player.Position.X, p.Player.Position.Y)
	screen.DrawImage(sprite, op)
}

func (p *PlayScene) drawCharacter(screen *ebiten.Image, c *Character) {
	if c == nil {
		return
	}
	gc := c.GetGameCharType()
	sprite := gc.GetSprite(c.GetAniIndex(), c.GetFaceDir())
	if sprite == nil {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.Position.X, c.Position.Y)
	screen.DrawImage(sprite, op)
}
func (p *PlayScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 128
}
