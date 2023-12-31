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

// YAML structure for config file
type ConfigFile struct {
	Triggers       []Trigger       `yaml:"triggers"`
	Chases         []Chase         `yaml:"chases"`
	EventSequences []EventSequence `yaml:"eventSequences"`
}

// A mapping of trigger source UID to trigger name
type Trigger struct {
	// Source, as key to map to a trigger
	Source string `yaml:"source" binding:"required"`
	// Goal of the trigger
	Goal string `yaml:"goal" binding:"required"`
	// Name of the target to work with, e.G. the name of a chase
	Target string `yaml:"target" binding:"required"`
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

// External representation of an event sequence (chain of triggered events).
type EventSequence struct {
	// Name for this chase - must be unique
	Name string `yaml:"name" binding:"required"`
	// The sequence of actions (chase) to take.
	Events []Event `yaml:"events" binding:"required"`
}

// External representation of an event sequence (chain of triggered events).
type Event struct {
	// Goal of the event
	Goal string `yaml:"goal" binding:"required"`
	// Name of the target to work with, e.G. the name of a chase
	Target string `yaml:"target" binding:"required"`
}
