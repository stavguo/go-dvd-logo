package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/stavguo/go-dvd-logo/core"
)

func RenderLogo(screen *ebiten.Image) {
	query := core.RenderFilter.Query()
	colors := core.ColorsResource.Get()
	// Iterate through entities with components.Position and components.Texture components
	for query.Next() {
		// Component access through the Query
		pos, texture := query.Get()

		if texture.Image != nil {
			// Get the current color
			currentColor := colors.Colors[texture.ColorIndex]

			// Create color matrix for colorm
			var cm colorm.ColorM
			cm.Scale(0, 0, 0, 0.9) // Reset RGB to 0, set alpha to 0.9

			// Set the new color
			r := float64(currentColor.R) / 255.0
			g := float64(currentColor.G) / 255.0
			b := float64(currentColor.B) / 255.0
			cm.Translate(r, g, b, 0)

			// Draw with color transformation
			op := &colorm.DrawImageOptions{}
			op.GeoM.Translate(pos.X, pos.Y)
			colorm.DrawImage(screen, texture.Image, cm, op)
		}
	}
}
