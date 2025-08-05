package systems

import (
	"github.com/stavguo/go-dvd-logo/core"
)

func CheckBounds() {
	query := core.BoundsFilter.Query()
	colors := core.ColorsResource.Get()
	screen := core.ScreenResource.Get()
	for query.Next() {
		pos, vel, texture := query.Get()

		if texture.Image != nil {
			bounds := texture.Image.Bounds()
			logoWidth := float64(bounds.Dx())
			logoHeight := float64(bounds.Dy())

			bounced := false

			// Check left and right bounds
			if pos.X <= 0 || pos.X+logoWidth >= float64(screen.Width) {
				vel.X = -vel.X
				bounced = true
			}

			// Check top and bottom bounds
			if pos.Y <= 0 || pos.Y+logoHeight >= float64(screen.Height) {
				vel.Y = -vel.Y
				bounced = true
			}

			// Change color when bounced
			if bounced {
				texture.ColorIndex = (texture.ColorIndex + 1) % len(colors.Colors)
			}
		}
	}
}
