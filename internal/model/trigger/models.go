package models_trigger

/*
A mapping of trigger source UID to trigger name
*/
type Trigger struct {
	// Source, as key to map to a trigger
	Source string
	// Name of the chase to work with
	Chase string
}
