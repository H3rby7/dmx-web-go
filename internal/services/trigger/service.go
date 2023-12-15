package trigger

import (
	models_event "github.com/H3rby7/dmx-web-go/internal/model/event"
	models_trigger "github.com/H3rby7/dmx-web-go/internal/model/trigger"
	"github.com/H3rby7/dmx-web-go/internal/services/config"
	"github.com/H3rby7/dmx-web-go/internal/services/event"
	log "github.com/sirupsen/logrus"
)

// Stateful construct to handle incoming triggers
type TriggerService struct {
	triggers             []models_trigger.Trigger
	eventSequenceService *event.EventSequenceService
}

// NewTriggerService creates a new [TriggerService] instance
//
// Also loads the triggers from the [ConfigService]
func NewTriggerService(configService *config.ConfigService, evSeqSvc *event.EventSequenceService) *TriggerService {
	log.Debugf("Creating new TriggerService")
	triggers := configService.GetTriggers()
	return &TriggerService{
		triggers:             triggers,
		eventSequenceService: evSeqSvc,
	}
}

// Handle the source by finding and running the corresponding trigger
func (svc *TriggerService) Handle(source string) bool {
	log.WithField("source", source).Infof("Handling incoming trigger")
	ok, trigger := svc.findTrigger(source)
	if !ok {
		return false
	}
	return svc.triggerTrigger(*trigger)
}

// findTrigger iterates over known triggers until it finds a [Trigger] with matching source.
//
// Returns a reference to the given trigger, if found.
// That way the same instance of the trigger is returned on censecutive calls with identical source
func (svc *TriggerService) findTrigger(source string) (ok bool, trigger *models_trigger.Trigger) {
	ll := log.WithField("source", source)

	ll.Debugf("Finding trigger")
	ok = false
	for i := 0; i < len(svc.triggers); i++ {
		if svc.triggers[i].Source == source {
			ok = true
			trigger = &svc.triggers[i]
			return
		}
	}
	ll.Warnf("No trigger with the given source")
	return
}

// triggerTrigger transforms the trigger into an event and passes it to the EventRouter
func (svc *TriggerService) triggerTrigger(trigger models_trigger.Trigger) bool {
	ev := models_event.Event{Goal: trigger.Goal, Target: trigger.Target}
	return svc.eventSequenceService.RouteEvent(ev)
}
