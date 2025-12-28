package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// PlayScene: simple gameplay placeholder moved to its own file
type PlayScene struct {
	sm         *SceneManager
	Player     *Player
	MapManager interface {
		Draw(screen *ebiten.Image)
		CanMoveHere(x, y float64) bool
	}
	Skeleton  *Character
	PlayingUI interface{ DrawUI(screen *ebiten.Image) }
	BtnExit   *CustomButton
}

func NewPlayScene(sm *SceneManager) *PlayScene {
	return &PlayScene{sm: sm, Player: NewPlayer()}

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
	// Render order similar to original Java: map, player, other chars, UI, buttons
	if p.MapManager != nil {
		p.MapManager.Draw(screen)
	}

	p.drawPlayer(screen)

	if p.PlayingUI != nil {
		p.PlayingUI.DrawUI(screen)
	}

	// Draw exit button if provided
	if p.BtnExit != nil {
		p.BtnExit.Draw(screen)
	}

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
