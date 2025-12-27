package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawTextAtCenter(screen *ebiten.Image, text string) {
	bounds := screen.Bounds()
	x := (bounds.Dx() - len(text)*7) / 2 // Approximate character width
	y := bounds.Dy() / 2
	ebitenutil.DebugPrintAt(screen, text, x, y)
}

func init() {
	var err error
	if err != nil {
		log.Fatal(err)
	}

}

type Game struct{}

func (g *Game) Update() error {
	PlayButtonNormal.SetPushed(false)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		PlayButtonNormal.SetPushed(true)
		log.Println("Button Pressed")
	} else {
		PlayButtonNormal.SetPushed(false)
		log.Println("Button Released")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawTextAtCenter(screen, "Bullet Quest 2D")
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		PlayButtonNormal.SetPushed(true)
		PlayButtonNormalPushed.Draw(screen)
		log.Println("Button Pressed")
	} else {
		PlayButtonNormal.SetPushed(false)
		PlayButtonNormal.Draw(screen)
		log.Println("Button Released")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 128
}

var PlayButtonNormal *CustomButton
var PlayButtonNormalPushed *CustomButton

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Bullet Quest 2D")
	buttonImageNormal := menuNormal
	buttonImagePushed := menuPushed
	PlayButtonNormal = NewCustomButton(100, 100, 64, 64, 1.0, buttonImageNormal)
	PlayButtonNormalPushed = NewCustomButton(100, 100, 64, 64, 1.0, buttonImagePushed)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

var (
	atlas, _, err = ebitenutil.NewImageFromFile("assets/bluebuttons.png")
	menuNormal    = atlas.SubImage(MenuNormalRect).(*ebiten.Image)
	menuPushed    = atlas.SubImage(MenuPushedRect).(*ebiten.Image) // if needed
	btn           = NewCustomButton(10, 10, 16, 16, 2.0, menuNormal)
)

func (g *Game) loadAssets() {
	if err != nil {
		log.Fatal(err)
	}

}
