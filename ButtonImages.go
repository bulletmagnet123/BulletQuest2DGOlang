package main

import (
	"image"
	"sync"
)

var MENU_START = image.Rect(0, 0, 320, 160)
var PLAYING_MENU = image.Rect(0, 0, 320, 160)

type ButtonImages int

const (
	MenuStart ButtonImages = iota
	PlayingMenu
)

var (
	// Playing frames: x=16, y=0, size=16x16
	PlayingNormalRect = image.Rect(16, 0, 16+16, 0+16)       // (16,0)-(32,16)
	PlayingPushedRect = image.Rect(16*11, 0, 16*11+16, 0+16) // (176,0)-(192,16)

	// Menu frames
	MenuNormalRect = image.Rect(48, 0, 48+16, 0+16)   // (48,0)-(64,16)
	MenuPushedRect = image.Rect(208, 0, 208+16, 0+16) // (208,0)-(224,16)
)

type ButtonState struct {
	PlayingNormal image.Image
	PlayingPushed image.Image
	MenuNormal    image.Image
	MenuPushed    image.Image
	Width         int
	Height        int
	ButtonHitbox  *image.Rectangle
}

var (
	buttonStates = make(map[ButtonImages]*ButtonState)
	mu           sync.RWMutex
)

func (b ButtonImages) Init(resID int, width, height int) {
	mu.Lock()
	defer mu.Unlock()

	state := &ButtonState{
		Width:  width,
		Height: height,
	}

	// Load button atlas and create sub-images
	// Note: You'll need to implement actual image loading based on your resource system
	// This is a placeholder structure

	buttonStates[b] = state
}

func (b ButtonImages) GetScaledBitmap(img image.Image, buttonScale int) image.Image {
	// Returns scaled bitmap - implementation depends on your image library
	return img
}

func (b ButtonImages) GetMenuButtonImage(isPushed bool) image.Image {
	mu.RLock()
	defer mu.RUnlock()

	state := buttonStates[b]
	if isPushed {
		return state.MenuPushed
	}
	return state.MenuNormal
}

func (b ButtonImages) GetPlayingButtonImage(isPushed bool) image.Image {
	mu.RLock()
	defer mu.RUnlock()

	state := buttonStates[b]
	if isPushed {
		return state.PlayingPushed
	}
	return state.PlayingNormal
}

func (b ButtonImages) getWidthPlayingNormal() int {
	mu.RLock()
	defer mu.RUnlock()
	state := buttonStates[b]
	return state.PlayingNormal.Bounds().Dx()
}

func (b ButtonImages) getHeightPlayingNormal() int {
	mu.RLock()
	defer mu.RUnlock()
	state := buttonStates[b]
	return state.PlayingNormal.Bounds().Dy()
}
func (b ButtonImages) getWidthPlayingPushed() int {
	mu.RLock()
	defer mu.RUnlock()
	state := buttonStates[b]
	return state.PlayingPushed.Bounds().Dx()
}
func (b ButtonImages) getHeightPlayingPushed() int {
	mu.RLock()
	defer mu.RUnlock()
	state := buttonStates[b]
	return state.PlayingPushed.Bounds().Dy()
}
func (b ButtonImages) getWidthMenuNormal() int {
	mu.RLock()
	defer mu.RUnlock()
	state := buttonStates[b]
	return state.MenuNormal.Bounds().Dx()
}
func (b ButtonImages) getHeightMenuNormal() int {
	mu.RLock()
	defer mu.RUnlock()
	state := buttonStates[b]
	return state.MenuNormal.Bounds().Dy()
}
func (b ButtonImages) getWidthMenuPushed() int {
	mu.RLock()
	defer mu.RUnlock()
	state := buttonStates[b]
	return state.MenuPushed.Bounds().Dx()
}
func (b ButtonImages) getHeightMenuPushed() int {
	mu.RLock()
	defer mu.RUnlock()
	state := buttonStates[b]
	return state.MenuPushed.Bounds().Dy()
}

func (b ButtonImages) Menu_getBtnImg(isBtnPushed bool) image.Image {
	if isBtnPushed {
		return b.GetPlayingButtonImage(true)
	}
	return b.GetPlayingButtonImage(false)
}

func (b ButtonImages) Playing_getBtnImg(isBtnPushed bool) image.Image {
	if isBtnPushed {
		return b.GetMenuButtonImage(true)
	}
	return b.GetMenuButtonImage(false)
}
