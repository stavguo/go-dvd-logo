package systems

import (
	"github.com/stavguo/go-dvd-logo/core"
)

func ApplyVelocity() {
	query := core.VelocityFilter.Query()
	for query.Next() {
		pos, vel := query.Get()
		pos.X += vel.X
		pos.Y += vel.Y
	}
}
