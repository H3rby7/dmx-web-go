// Package models_config contains the structures used by the [config] package
//
// Provides the structures including the YAML bindings, to read information from the filesystem
package models_config

//	Struct holding a channel and a value for that channel.
type DMXValueForChannel struct {
	Channel int16 `yaml:"channel" binding:"required"`
	Value   byte  `yaml:"value" binding:"required"`
}

//	Struct holding a list of channels and their values
type Scene struct {
	List []DMXValueForChannel `yaml:"list"`
}

// File containing the chases and their triggers
type ConfigFile struct {
	Triggers []Trigger `yaml:"triggers"`
	Chases   []Chase   `yaml:"chases"`
}

// A mapping of trigger source UID to trigger name
type Trigger struct {
	// Source, as key to map to a trigger
	Source string `yaml:"source" binding:"required"`
	// Name of the chase to work with
	Chase string `yaml:"chase" binding:"required"`
}

// External representation of a Chase (chain of actions).
type Chase struct {
	// Name for this chase - must be unique
	Name string `yaml:"name" binding:"required"`
	// The sequence of actions (chase) to take.
	Chase []Step `yaml:"chase" binding:"required"`
}

// An element in the list of a chase
type Step struct {

	// Time in millis to pass after the previous item before this scene is faded in.
	//
	// Default: 0
	DelayTimeMillis int64 `yaml:"delayTimeMillis"`

	// 	The desired status of the DMX bridge.
	//
	// 	Default: unset, no changes done.
	BridgeActive bool `yaml:"bridgeActive"`

	// Scene content to transition to
	//
	// Default: no scene
	Scene Scene `yaml:"scene"`

	// Fade time in millis for the change.
	//
	// Default = 0
	//
	// *Note: The bridge does not support being faded in/out yet.*
	FadeTimeMillis int64 `yaml:"fadeTimeMillis"`
}
