package models_chase

import (
	models_scene "github.com/H3rby7/dmx-web-go/internal/model/scene"
)

// An element in the list of a chase
type Step struct {
	/*
		Time in millis to pass after the previous item before this scene is faded in.

		Default: 0
	*/
	DelayTimeMillis int64
	/*
		The desired status of the DMX bridge.

		Default: unset, no changes done.
	*/
	BridgeActive bool
	/*
		Scene content to transition to

		Default: no scene
	*/
	Scene models_scene.Scene
	/*
		Fade time in millis for the change.

		Default = 0

		*Note: The bridge does not support being faded in/out yet.*
	*/
	FadeTimeMillis int64
}
