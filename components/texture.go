package components

import "github.com/hajimehoshi/ebiten/v2"

type Texture struct {
	Image      *ebiten.Image
	ColorIndex int
}
