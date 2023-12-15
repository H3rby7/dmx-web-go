// Package models_event defines [Event] and [EventSequence] as internal structures.
//
// These structures are used to create stateful trigger behaviour.
//
// Also an [models_trigger.Trigger] will be mapped to an event for architectural reasons.
package models_event

import log "github.com/sirupsen/logrus"

// EventDelegateFunc defines the function to handle events as they are fired.
type EventDelegateFunc func(event Event) bool

// Internal representation of an EventSequence.
type EventSequence struct {
	// Name for this sequence - must be unique
	Name string
	// The sequence of events to trigger upon invoking.
	Events []Event
	// Index of next event (in the sequence)
	nextEvent int
}

// Internal representation of an Event (goal+target)
//
// Item in the [EventSequence]
type Event struct {
	// Goal of the event
	Goal string
	// Name of the target to work with, e.G. the name of a chase
	Target string
}

// NewEventSequence constructs a EventSequence struct, given the name and events to use.
func NewEventSequence(name string, events []Event) EventSequence {
	es := EventSequence{
		Name:      name,
		Events:    events,
		nextEvent: 0,
	}
	return es
}

// Fire the next event and move forward in the sequence
//
// When the end of the sequence is reached, resets to start
func (s *EventSequence) FireNextEvent(delegate EventDelegateFunc) bool {
	e := s.Events[s.nextEvent]
	log.Infof("Next event (%d) has goal '%s' with target '%s'", s.nextEvent, e.Goal, e.Target)
	s.nextEvent = (s.nextEvent + 1) % len(s.Events)
	log.Tracef("Next event will be '%d", s.nextEvent)
	return delegate(e)
}
