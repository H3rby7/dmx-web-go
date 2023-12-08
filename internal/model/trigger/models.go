// Package models_trigger provides the internal representations of [Trigger]s
package models_trigger

// A mapping of trigger source UID to trigger name
type Trigger struct {
	// Source, as key to map to a trigger
	Source string
	// Goal of the trigger
	Goal string
	// Name of the target to work with, e.G. the name of a chase
	Target string
}
