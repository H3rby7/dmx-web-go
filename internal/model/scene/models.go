package models_scene

import models_dmx "github.com/H3rby7/dmx-web-go/internal/model/dmx"

/*
	Struct holding a list of channels and their values
*/
type Scene struct {
	List []models_dmx.DMXValueForChannel
}

/*
	Struct holding:

	* a list of channels and their values

	* fade time in millis to reach the given channels and values
*/
type SceneWithFade struct {
	Scene          Scene
	FadeTimeMillis int64
}
