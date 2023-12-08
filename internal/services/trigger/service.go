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

func NewTriggerService(configService *config.ConfigService, chaseService *chase.ChaseService) *TriggerService {
	log.Debugf("Creating new TriggerService")
	triggers := configService.GetTriggers()
	return &TriggerService{
		triggers:     triggers,
		chaseService: chaseService,
	}
}

// Handle an incoming trigger
func (svc *TriggerService) Handle(source string) (ok bool) {
	log.WithField("source", source).Infof("Handling incoming trigger")
	ok, chase := svc.mapToChaseName(source)
	if ok {
		// TODO: Handle triggered again, before chase has finished
		// TODO: Other triggers, like start/stop/continue etc.
		svc.chaseService.StartChaseFromTheTop(chase)
	}
	return
}

func (svc *TriggerService) mapToChaseName(source string) (ok bool, chaseName string) {
	ll := log.WithField("source", source)

	ll.Debugf("Mapping trigger source to chase")
	ok = false
	for _, m := range svc.triggers {
		if m.Source == source {
			ok = true
			chaseName = m.Chase
			ll.Debugf("Matched trigger to chase with name '%s'", chaseName)
			return
		}
	}
	ll.Warnf("No trigger with the given source")
	return
}
