package assets

import (
	_ "embed"
)

//go:embed Cabin-Bold.ttf
var CabinBoldFontData []byte

//go:embed Cabin-Regular.ttf
var CabinRegularFontData []byte

//go:embed logo.png
var LogoPNGData []byte
