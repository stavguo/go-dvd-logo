package core

import (
	"github.com/mlange-42/ark/ecs"
)

var (
	World *ecs.World
)

func InitWorld() {
	world := ecs.NewWorld()
	World = &world
}
