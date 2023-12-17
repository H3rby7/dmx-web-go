// Package models_dmx contains the internal representations DMX Channels and their values
package models_dmx

// Struct holding a channel and a value for that channel.
type DMXValueForChannel struct {
	Channel int16
	Value   byte
}

// Struct holding a list of channels and their values
type MultipleDMXValueForChannel struct {
	List []DMXValueForChannel
}
