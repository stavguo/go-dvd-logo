package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stavguo/go-dvd-logo/scenes"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	sceneManager *scenes.SceneManager
}

func (g *Game) Update() error {
	return g.sceneManager.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	sceneManager := scenes.NewSceneManager()
	titleScene := scenes.NewTitleScene(sceneManager)
	sceneManager.SetScene(titleScene)

	// gameScene := NewGameScene()
	// sceneManager.SetScene(gameScene)

	game := &Game{sceneManager: sceneManager}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Bouncing DVD Logo")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
