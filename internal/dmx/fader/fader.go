package dmxfader

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type DMXFader struct {
	isActive     bool
	channel      int16
	currentValue float32
	targetValue  byte
	deadline     time.Time
}

/*
	Give the fader a new target and deadline

To switch a value immediately simply pass '0' as fadeTimeMillis
*/
func (f *DMXFader) FadeTo(targetValue byte, fadeTimeMillis int64) {
	f.targetValue = targetValue
	f.deadline = time.Now().Add(time.Millisecond * time.Duration(fadeTimeMillis))
	f.isActive = true
	log.Debugf("Fading channel '%v' to '%v' until %v", f.channel, targetValue, f.deadline)
}

// Calculate the next DMX value
func (f *DMXFader) GetNextValue() byte {
	deltaT := f.deadline.UnixMilli() - time.Now().UnixMilli()
	if deltaT < 0 {
		f.isActive = false
		return f.targetValue
	}
	stepsTillDeadline := deltaT / TICK_INTERVAL_MILLIS
	deltaV := float32(f.targetValue) - f.currentValue
	nextValue := byte(f.currentValue + deltaV/float32(stepsTillDeadline))
	log.Tracef("next value for channel '%d' is '%d'", f.channel, nextValue)
	return nextValue
}

// Is this fader active, or already done?
func (f *DMXFader) IsActive() bool {
	return f.isActive
}
