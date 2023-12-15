// Package apiv1dtos defines the models and their JSON bindings for the API in Version 1
package apiv1dtos

// DMXValueForChannel describes a channel and a value for it.
type DMXValueForChannel struct {
	Channel int16 `json:"channel" binding:"required"`
	Value   byte  `json:"value" binding:"required"`
}

// MultipleDMXValueForChannel holds a list of channels and their values
type MultipleDMXValueForChannel struct {
	List []DMXValueForChannel `json:"list" binding:"required"`
}

// Alias for MultipleDMXValueForChannel
type Scene = MultipleDMXValueForChannel

// MultipleDMXValueForChannelWithFade holds:
//
// * a list of channels and their values
//
// * fade time in millis to reach the given channels and values
type MultipleDMXValueForChannelWithFade struct {
	Scene          Scene `json:"scene" binding:"required"`
	FadeTimeMillis int64 `json:"fadeTimeMillis" binding:"required"`
}

// SceneWithFade is an alias for MultipleDMXValueForChannelWithFade
type SceneWithFade = MultipleDMXValueForChannelWithFade

// TriggerSource simply holds the source of the request
type TriggerSource struct {
	Source string `json:"source"`
}
