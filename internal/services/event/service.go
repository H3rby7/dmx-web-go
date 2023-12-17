// Package event hosts the EventService and provides means to work with Events and Event Sequences
package event

import (
	models_event "github.com/H3rby7/dmx-web-go/internal/model/event"
	"github.com/H3rby7/dmx-web-go/internal/services/chase"
	"github.com/H3rby7/dmx-web-go/internal/services/config"
	log "github.com/sirupsen/logrus"
)

// EventService handles chases and actions on chases
type EventService struct {
	sequences    []models_event.EventSequence
	chaseService *chase.ChaseService
}

// NewEventService creates a new [EventService] instance
//
// Also loads the event sequences from the [ConfigService]
func NewEventService(configService *config.ConfigService, chaseService *chase.ChaseService) *EventService {
	log.Debugf("Creating new EventService")
	sequences := configService.GetEventSequences()
	return &EventService{
		sequences:    sequences,
		chaseService: chaseService,
	}
}

// FireNextEventFromSequence fires the next event in the sequence
func (svc *EventService) FireNextEventFromSequence(sequenceName string) (ok bool) {
	ll := log.WithField("sequence", sequenceName)
	ok, sequence := svc.findSequenceByName(sequenceName)
	if !ok {
		ll.Warnf("Could not find event sequence")
		return
	}
	ll.Debugf("Found event sequence")
	return sequence.FireNextEvent(svc.RouteEvent)
}

// findSequenceByName looks for a sequence matching the name
//
// Returns a reference to the given sequence, if found.
// That way the same instance of the sequence is returned on censecutive calls with identical sequenceName
func (svc *EventService) findSequenceByName(sequenceName string) (ok bool, seq *models_event.EventSequence) {
	ok = false
	for i := 0; i < len(svc.sequences); i++ {
		if svc.sequences[i].Name == sequenceName {
			ok = true
			seq = &svc.sequences[i]
			return
		}
	}
	return
}

// RouteEvent runs an action depending on the goal
func (svc *EventService) RouteEvent(event models_event.Event) bool {
	ll := log.WithField("goal", event.Goal).WithField("target", event.Target)
	switch event.Goal {
	case "chase-reset-and-run":
		return svc.chaseService.StartChaseFromTheTop(event.Target)
	case "event-sequence-next":
		return svc.FireNextEventFromSequence(event.Target)
	default:
		ll.Warnf("Goal '%s' is unknown", event.Goal)
		return false
	}
}
