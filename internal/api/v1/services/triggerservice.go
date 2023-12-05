package apiv1services

import (
	"github.com/H3rby7/dmx-web-go/internal/trigger"
	log "github.com/sirupsen/logrus"
)

// Stateful construct to handle incoming triggers
type TriggerService struct {
	triggers []trigger.Trigger
	chases   []trigger.Chase
}

func NewTriggerService() *TriggerService {
	log.Debugf("Creating new TriggerService")
	actions := trigger.Load()
	log.Infof("Loaded %d triggers for %d chases", len(actions.Triggers), len(actions.Chases))
	return &TriggerService{
		triggers: actions.Triggers,
		chases:   actions.Chases,
	}
}

/*
Handle an incoming trigger
*/
func (svc *TriggerService) Handle(source string) (ok bool) {
	ll := log.WithField("source", source)
	ll.Infof("Handling incoming trigger")
	ok, trigger := svc.mapToChase(source)
	if !ok {
		return ok
	}
	// TODO: Pass to Chase Service
	ll.Warnf("Work-In-Progress: Applying only first scene of chase!")
	wip := trigger.Chase[0]
	SetScene(wip.Scene, wip.FadeTimeMillis)
	return
}

func (svc *TriggerService) mapToChase(source string) (ok bool, c trigger.Chase) {
	ll := log.WithField("source", source)

	ll.Debugf("Mapping to trigger")
	ok = false
	chaseName := ""
	for _, m := range svc.triggers {
		if m.Source == source {
			ok = true
			chaseName = m.Chase
			ll.Debugf("Matched trigger with name '%s'", source)
			break
		}
	}
	if !ok {
		ll.Warnf("Cannot be mapped to any trigger")
		return
	}
	ok = false
	for _, chase := range svc.chases {
		if chase.Name == chaseName {
			ll.Debugf("Returning chase with name '%s'", chaseName)
			ok = true
			c = chase
			return
		}
	}
	ll.Warnf("Could not find chase with name '%s'", chaseName)
	return
}
