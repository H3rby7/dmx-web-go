// Package models_fader defines [DMXFader] as a central piece of the DMX-Web architecture.
//
// The [DMXFader] takes care of the value of one DMX channel and allows fading the value.
package models_fader

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// TICK_INTERVAL_MILLIS can also be seen as the fading framerate.
const TICK_INTERVAL_MILLIS = 1000 / 25 // 25 updates per second

// A value of 1 millisecond results in immediate fading
const FADE_IMMEDIATELY int64 = 1

// NewDMXFader creates a new [DMXFader] instance with defaults for the given channel
func NewDMXFader(channel int16) DMXFader {
	return DMXFader{
		isActive: false,
		channel:  channel,
	}
}

// DMXFader is a struct following the operator pattern.
//
// At any given time the struct aims to achieve the targetValue within the deadline.
//
// It does so by adjusting the currentValue in a linear fashion towards the targetValue
type DMXFader struct {
	// Is this fader working to reach a new value?
	isActive bool
	// The channel connected to this fader
	channel int16
	// Current DMX value of the fader
	currentValue float32
	// Target DMX value for the fader
	targetValue byte
	// Timestamp when the targetValue will be reached
	deadline time.Time
}

// Give the fader a new target and deadline
//
// To switch a value immediately simply pass '0' as fadeTimeMillis
func (f *DMXFader) FadeTo(targetValue byte, fadeTimeMillis int64) {
	// Check if the fader is not fading right now and already has the target value
	if !f.isActive && (byte(f.currentValue) == targetValue) {
		log.Tracef("Skipping channel '%v', as it already rests at target value '%v'", f.channel, targetValue)
		return
	}
	f.targetValue = targetValue
	f.deadline = time.Now().Add(time.Millisecond * time.Duration(fadeTimeMillis))
	f.isActive = true
	log.Debugf("Fading channel '%v' from '%v' to '%v' until %v", f.channel, f.currentValue, f.targetValue, f.deadline)
}

// Calculates and updates internal state
//
// Returns the new value for convenience
func (f *DMXFader) UpdateValue() byte {
	deltaT := f.deadline.UnixMilli() - time.Now().UnixMilli()
	if deltaT < TICK_INTERVAL_MILLIS {
		// We reached the end of our fade
		f.isActive = false
		f.currentValue = float32(f.targetValue)
		log.Debugf("Fader for channel '%d' reached target '%d'", f.channel, f.targetValue)
		// Return target value for convenience
		return f.GetCurrentValue()
	}
	// Iterations left, before reaching the deadline
	stepsTillDeadline := deltaT / TICK_INTERVAL_MILLIS
	// Total value difference to cover
	deltaV := float32(f.targetValue) - f.currentValue
	// Value change in this iteration
	update := deltaV / float32(stepsTillDeadline)
	// New value
	f.currentValue = f.currentValue + update
	log.Tracef("Next value for channel '%d' is '%.1f'", f.channel, f.currentValue)
	// Return new value for convenience
	return f.GetCurrentValue()
}

// Returns the current value
func (f *DMXFader) GetCurrentValue() byte {
	return byte(f.currentValue)
}

// Is this fader active, or already done?
func (f *DMXFader) IsActive() bool {
	return f.isActive
}

// Set internal value
//
// Use sparsely!
func (f *DMXFader) SetValue(val float32) {
	f.currentValue = val
}
