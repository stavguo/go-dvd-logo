package core

import (
	"github.com/mlange-42/ark/ecs"
	"github.com/stavguo/go-dvd-logo/components"
)

var (
	ColorsResource *ecs.Resource[components.Colors]
	ScreenResource *ecs.Resource[components.Screen]
)

func InitResources() {
	colRes := ecs.NewResource[components.Colors](World)
	ColorsResource = &colRes

	screenRes := ecs.NewResource[components.Screen](World)
	ScreenResource = &screenRes
}
