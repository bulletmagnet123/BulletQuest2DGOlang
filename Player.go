package main

// Converted from:
// public class Player extends Character {
//     public Player() {
//         super(new PointF(GAME_WIDTH / 2, GAME_HEIGHT / 2), GameCharacters.PLAYER);
//     }
//
//     public void update(double delta, boolean movePlayer) {
//         if (movePlayer)
//             updateAnimation();
//     }
// }

// Minimal supporting types (adjust or remove if you have your own implementations)
type PointF struct {
	X float64
	Y float64
}

type GameCharacter int

const (
	GameCharacterPlayer GameCharacter = iota
)

const (
	GAME_WIDTH  = 320
	GAME_HEIGHT = 128
)

type Character struct {
	Position     PointF
	GameCharType GameCharacter
	AniTick      int
	AniIndex     int
	FaceDir      int
}

const (
	FACE_DIR_DOWN  = 0
	FACE_DIR_UP    = 1
	FACE_DIR_LEFT  = 2
	FACE_DIR_RIGHT = 3

	ANIM_DEFAULT_SPEED = 10
	ANIM_AMOUNT        = 4
)

func NewCharacter(pos PointF, t GameCharacter) *Character {
	return &Character{
		Position:     pos,
		GameCharType: t,
		AniTick:      0,
		AniIndex:     0,
		FaceDir:      FACE_DIR_DOWN,
	}
}

func (c *Character) UpdateAnimation() {
	c.AniTick++
	if c.AniTick >= ANIM_DEFAULT_SPEED {
		c.AniTick = 0
		c.AniIndex++
		if c.AniIndex >= ANIM_AMOUNT {
			c.AniIndex = 0
		}
	}
}

func (c *Character) ResetAnimation() {
	c.AniTick = 0
	c.AniIndex = 0
}

func (c *Character) GetAniIndex() int {
	return c.AniIndex
}

func (c *Character) GetFaceDir() int {
	return c.FaceDir
}

func (c *Character) SetFaceDir(faceDir int) {
	c.FaceDir = faceDir
}

func (c *Character) GetGameCharType() GameCharacter {
	return c.GameCharType
}

// Player in Go

type Player struct {
	*Character
	Attacking bool
}

func NewPlayer() *Player {
	return &Player{NewCharacter(PointF{X: GAME_WIDTH / 2, Y: GAME_HEIGHT / 2}, GameCharacterPlayer), false}
}

func (p *Player) Update(delta float64, movePlayer bool) {
	if movePlayer {
		p.UpdateAnimation()
	}
}
