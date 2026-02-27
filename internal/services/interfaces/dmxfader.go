// Package interfaces defines service interfaces to interact with DMX
package interfaces

import (
	models_scene "github.com/H3rby7/dmx-web-go/internal/model/scene"
)

// DMX Writer that takes care of fading channels to the desired values over time.
type DMXFader interface {
	ClearAll()
	FadeScene(scene models_scene.Scene, fadeDurationMillis int64)
	FadeTo(channel int16, value byte, fadeDurationMillis int64)
	Start()
	Stop()
	CleanUp()
}
