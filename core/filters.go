package core

import (
	"github.com/mlange-42/ark/ecs"
	"github.com/stavguo/go-dvd-logo/components"
)

var (
	RenderFilter   *ecs.Filter2[components.Position, components.Texture]
	VelocityFilter *ecs.Filter2[components.Position, components.Velocity]
	BoundsFilter   *ecs.Filter3[components.Position, components.Velocity, components.Texture]
)

func InitFilters() {
	RenderFilter = ecs.NewFilter2[components.Position, components.Texture](World)
	VelocityFilter = ecs.NewFilter2[components.Position, components.Velocity](World)
	BoundsFilter = ecs.NewFilter3[components.Position, components.Velocity, components.Texture](World)
}
