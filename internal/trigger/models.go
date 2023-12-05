package trigger

import apiv1models "github.com/H3rby7/dmx-web-go/internal/api/v1/models"

/*
	File containing the chases and their triggers
*/
type ActionsFile struct {
	Triggers []Trigger `yaml:"triggers" binding:"required"`
	Chases   []Chase   `yaml:"chases" binding:"required"`
}

/*
	A mapping of trigger source UID to trigger name
*/
type Trigger struct {
	// Source, as key to map to a trigger
	Source string `yaml:"source" binding:"required"`
	// Name of the chase to work with
	Chase string `yaml:"chase" binding:"required"`
}

/*
	A Chase (chain of actions).
*/
type Chase struct {
	// Name for this trigger - mus be unique
	Name string `yaml:"name" binding:"required"`
	// The sequence of actions (chase) to take.
	Chase []ChaseElement `yaml:"chase" binding:"required"`
}

// An element in the list of a chase
type ChaseElement struct {
	/*
		Time in millis to pass after the previous item before this scene is faded in.

		Default: 0
	*/
	DelayTimeMillis int64 `yaml:"delayTimeMillis"`
	/*
		The desired status of the DMX bridge.

		Default: unset, no changes done.
	*/
	BridgeActive bool `yaml:"bridgeActive"`
	/*
		Scene content to transition to

		Default: no scene
	*/
	Scene apiv1models.Scene `yaml:"scene"`
	/*
		Fade time in millis for the change.

		Default = 0

		*Note: The bridge does not support being faded in/out yet.*
	*/
	FadeTimeMillis int64 `yaml:"fadeTimeMillis"`
}
