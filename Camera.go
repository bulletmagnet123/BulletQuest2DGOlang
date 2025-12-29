package main

type Camera struct {
	X, Y       float64
	ScreenW    int
	ScreenH    int
	WorldW     int
	WorldH     int
	FollowedCh *Character
}

func NewCamera(screenW, screenH, worldW, worldH int) *Camera {
	return &Camera{
		ScreenW: screenW,
		ScreenH: screenH,
		WorldW:  worldW,
		WorldH:  worldH,
	}
}

func (c *Camera) FollowCharacter(ch *Character) {
	c.FollowedCh = ch
}

func (c *Camera) Update() {
	if c.FollowedCh == nil {
		return
	}

	// Center the camera on the followed character
	// Character doesn't currently expose Width/Height, so follow its position directly.
	c.X = c.FollowedCh.Position.X - float64(c.ScreenW)/2
	c.Y = c.FollowedCh.Position.Y - float64(c.ScreenH)/2

	// Clamp the camera position to the world bounds
	if c.X < 0 {
		c.X = 0
	}
	if c.Y < 0 {
		c.Y = 0
	}
	if c.X > float64(c.WorldW-c.ScreenW) {
		c.X = float64(c.WorldW - c.ScreenW)
	}
	if c.Y > float64(c.WorldH-c.ScreenH) {
		c.Y = float64(c.WorldH - c.ScreenH)
	}

}
