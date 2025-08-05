package scenes

import "github.com/hajimehoshi/ebiten/v2"

// Scene interface that all scenes must implement
type Scene interface {
	Enter()
	Update() error
	Draw(screen *ebiten.Image)
	Exit()
}

type SceneManager struct {
	currentScene Scene
	nextScene    Scene
	shouldChange bool
}

func NewSceneManager() *SceneManager {
	return &SceneManager{}
}

func (sm *SceneManager) SetScene(scene Scene) {
	sm.nextScene = scene
	sm.shouldChange = true
}

// Update handles scene transitions and updates current scene
func (sm *SceneManager) Update() error {
	// Handle scene transition
	if sm.shouldChange {
		if sm.currentScene != nil {
			sm.currentScene.Exit()
		}
		sm.currentScene = sm.nextScene
		sm.nextScene = nil
		sm.shouldChange = false
		if sm.currentScene != nil {
			sm.currentScene.Enter()
		}
	}

	// Update current scene
	if sm.currentScene != nil {
		return sm.currentScene.Update()
	}
	return nil
}

func (sm *SceneManager) Draw(screen *ebiten.Image) {
	if sm.currentScene != nil {
		sm.currentScene.Draw(screen)
	}
}
