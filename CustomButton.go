package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type CustomButton struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
	Scale  float64
	Pushed bool
	Image  *ebiten.Image
}

func NewCustomButton(x, y, width, height, scale float64, img *ebiten.Image) *CustomButton {
	return &CustomButton{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Scale:  scale,
		Image:  img,
	}
}

func (b *CustomButton) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(b.Scale, b.Scale)
	op.GeoM.Translate(b.X, b.Y)
	screen.DrawImage(b.Image, op) // draw the sub-image directly
}

func (b *CustomButton) SetPushed(pushed bool) {
	b.Pushed = pushed
}

func (b *CustomButton) IsPushed() bool {
	return b.Pushed
}

func (b *CustomButton) Contains(mx, my int) bool {
	left := int(b.X)
	top := int(b.Y)
	w := int(b.Width * b.Scale)
	h := int(b.Height * b.Scale)
	return mx >= left && mx < left+w && my >= top && my < top+h
}

func loadImage() *ebiten.Image {
	atlas, _, err := ebitenutil.NewImageFromFile("assets/bluebuttons.png")
	if err != nil {
		log.Fatal(err)
	}

	menuNormal := atlas.SubImage(MenuNormalRect).(*ebiten.Image)
	_ = atlas.SubImage(MenuPushedRect).(*ebiten.Image) // if needed

	return menuNormal
}
