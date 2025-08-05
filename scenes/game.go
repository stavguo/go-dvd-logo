package scenes

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand/v2"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/ark/ecs"
	"github.com/stavguo/go-dvd-logo/components"
	"github.com/stavguo/go-dvd-logo/core"
	"github.com/stavguo/go-dvd-logo/systems"
)

const (
	speed = 1.0
)

type GameScene struct{}

func NewGameScene() *GameScene {
	return &GameScene{}
}

func (gs *GameScene) Enter() {
	// Initialize core systems
	core.InitWorld()
	core.InitFilters()
	core.InitResources()

	core.ScreenResource.Add(&components.Screen{Width: 640, Height: 480})

	// Add the colors resource using the resource mapper
	core.ColorsResource.Add(&components.Colors{Colors: []color.RGBA{
		{225, 6, 4, 255},   // Red
		{0, 238, 0, 255},   // Vivid Green
		{0, 0, 255, 255},   // Blue
		{0, 255, 255, 255}, // Cyan
		{255, 0, 255, 255}, // Magenta
		{255, 255, 0, 255}, // Yellow
		{255, 155, 0, 255}, // Orange
	}})

	// Create a component mapper for all 3 components
	mapper := ecs.NewMap3[components.Position, components.Velocity, components.Texture](core.World)

	// Load the PNG first to get its dimensions
	logoImg := gs.loadPNG()
	bounds := logoImg.Bounds()
	logoWidth := float64(bounds.Dx())
	logoHeight := float64(bounds.Dy())

	// Calculate valid spawn area s.t. the logo won't clip out of bounds
	screen := core.ScreenResource.Get()
	maxX := screen.Width - int(logoWidth)
	maxY := screen.Height - int(logoHeight)

	randomX := float64(rand.IntN(maxX))
	randomY := float64(rand.IntN(maxY))

	// Randomly choose direction for X and Y
	velX := speed
	if rand.IntN(2) == 0 {
		velX = -speed
	}

	velY := speed
	if rand.IntN(2) == 0 {
		velY = -speed
	}

	_ = mapper.NewEntity(
		&components.Position{X: randomX, Y: randomY},
		&components.Velocity{X: velX, Y: velY},
		&components.Texture{Image: logoImg, ColorIndex: rand.IntN(len(core.ColorsResource.Get().Colors))},
	)
}

func (gs *GameScene) loadPNG() *ebiten.Image {
	// Read the PNG file
	pngData, err := os.ReadFile("assets/logo.png")
	if err != nil {
		log.Fatal("Failed to read PNG file:", err)
	}

	// Decode the PNG image
	img, _, err := image.Decode(bytes.NewReader(pngData))
	if err != nil {
		log.Fatal("Failed to decode PNG:", err)
	}

	// Convert to Ebiten image
	return ebiten.NewImageFromImage(img)
}

func (gs *GameScene) Update() error {
	systems.ApplyVelocity()
	systems.CheckBounds()
	return nil
}

func (gs *GameScene) Draw(screen *ebiten.Image) {
	systems.RenderLogo(screen)
}

func (gs *GameScene) Exit() {}
