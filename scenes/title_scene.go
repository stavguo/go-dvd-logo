package scenes

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/stavguo/go-dvd-logo/assets"
)

var (
	cabinBoldFontSource    *text.GoTextFaceSource
	cabinRegularFontSource *text.GoTextFaceSource
)

// loadFontSourceFromBytes creates a GoTextFaceSource from embedded font data
func loadFontSourceFromBytes(fontData []byte) (*text.GoTextFaceSource, error) {
	source, err := text.NewGoTextFaceSource(bytes.NewReader(fontData))
	if err != nil {
		return nil, err
	}
	return source, nil
}

func init() {
	var err error

	// Load Cabin-Bold font from embedded data
	cabinBoldFontSource, err = loadFontSourceFromBytes(assets.CabinBoldFontData)
	if err != nil {
		log.Fatal("Failed to load Cabin-Bold font:", err)
	}

	// Load Cabin-Regular font from embedded data
	cabinRegularFontSource, err = loadFontSourceFromBytes(assets.CabinRegularFontData)
	if err != nil {
		log.Fatal("Failed to load Cabin-Regular font:", err)
	}
}

type TitleScene struct {
	sceneManager *SceneManager
}

// NewTitleScene creates a new title scene
func NewTitleScene(sceneManager *SceneManager) *TitleScene {
	return &TitleScene{sceneManager: sceneManager}
}

// Enter is called when entering this scene
func (ts *TitleScene) Enter() {
	// Title scene initialization if needed
}

func (ts *TitleScene) Update() error {
	// Check for any key press
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) ||
		inpututil.IsKeyJustPressed(ebiten.KeyEnter) ||
		len(inpututil.AppendJustPressedKeys(nil)) > 0 {
		// Switch to game scene
		gameScene := NewGameScene()
		ts.sceneManager.SetScene(gameScene)
	}
	return nil
}

func (ts *TitleScene) Draw(screen *ebiten.Image) {

	const (
		titleFontSize    = 48
		subtitleFontSize = 20
	)

	// Draw title text
	titleText := "Bouncing DVD Logo Demo"
	titleOp := &text.DrawOptions{}

	// Center the title text
	titleWidth, _ := text.Measure(titleText, &text.GoTextFace{
		Source: cabinBoldFontSource,
		Size:   titleFontSize,
	}, 0)
	titleX := (640 - titleWidth) / 2
	titleY := 200

	titleOp.GeoM.Translate(float64(titleX), float64(titleY))
	titleOp.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 255})
	text.Draw(screen, titleText, &text.GoTextFace{
		Source: cabinBoldFontSource,
		Size:   titleFontSize,
	}, titleOp)

	// Draw subtitle text
	subtitleText := "Press any key to continue"
	subtitleOp := &text.DrawOptions{}

	// Center the subtitle text
	subtitleWidth, _ := text.Measure(subtitleText, &text.GoTextFace{
		Source: cabinRegularFontSource,
		Size:   subtitleFontSize,
	}, 0)
	subtitleX := (640 - subtitleWidth) / 2
	subtitleY := 300

	subtitleOp.GeoM.Translate(float64(subtitleX), float64(subtitleY))
	subtitleOp.ColorScale.ScaleWithColor(color.RGBA{200, 200, 200, 255})
	text.Draw(screen, subtitleText, &text.GoTextFace{
		Source: cabinRegularFontSource,
		Size:   subtitleFontSize,
	}, subtitleOp)
}

func (ts *TitleScene) Exit() {}
