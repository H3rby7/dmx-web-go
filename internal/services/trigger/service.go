package trigger

import (
	models_trigger "github.com/H3rby7/dmx-web-go/internal/model/trigger"
	"github.com/H3rby7/dmx-web-go/internal/services/chase"
	"github.com/H3rby7/dmx-web-go/internal/services/config"
	log "github.com/sirupsen/logrus"
)

// Stateful construct to handle incoming triggers
type TriggerService struct {
	triggers     []models_trigger.Trigger
	chaseService *chase.ChaseService
}

// NewTriggerService creates a new [TriggerService] instance
//
// Also loads the triggers from the [ConfigService]
func NewTriggerService(configService *config.ConfigService, chaseService *chase.ChaseService) *TriggerService {
	log.Debugf("Creating new TriggerService")
	triggers := configService.GetTriggers()
	return &TriggerService{
		triggers:     triggers,
		chaseService: chaseService,
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

// triggerTrigger runs an action depending on the goal
func (svc *TriggerService) triggerTrigger(trigger models_trigger.Trigger) bool {
	ll := log.WithField("goal", trigger.Goal).WithField("target", trigger.Target)
	switch trigger.Goal {
	// TODO: Other triggers, like start/stop/continue etc.
	case "chase-reset-and-run":
		svc.chaseService.StartChaseFromTheTop(trigger.Target)
	default:
		ll.Warnf("Goal '%s' is unknown", trigger.Goal)
		return false
	}
	return true
}
