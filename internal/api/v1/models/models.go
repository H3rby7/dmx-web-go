package apiv1models

/*
	Struct holding a channel and a value for that channel.
*/
type DMXValueForChannel struct {
	Channel int16 `json:"channel" binding:"required"`
	Value   byte  `json:"value" binding:"required"`
}

/*
	Struct holding a list of channels and their values
*/
type MultipleDMXValueForChannel struct {
	List []DMXValueForChannel `json:"list" binding:"required"`
}

// Alias for MultipleDMXValueForChannel
type Scene = MultipleDMXValueForChannel

/*
	Struct holding:

	* a list of channels and their values

	* fade time in millis to reach the given channels and values
*/
type MultipleDMXValueForChannelWithFade struct {
	Scene          Scene `json:"scene" binding:"required"`
	FadeTimeMillis int64 `json:"fadeTimeMillis" binding:"required"`
}

// Alias for MultipleDMXValueForChannelWithFade
type SceneWithFade = MultipleDMXValueForChannelWithFade

type TriggerSource struct {
	Source string `json:"source"`
}
