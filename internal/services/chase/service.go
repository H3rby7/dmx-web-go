package chase

import (
	models_chase "github.com/H3rby7/dmx-web-go/internal/model/chase"
	models_scene "github.com/H3rby7/dmx-web-go/internal/model/scene"
	"github.com/H3rby7/dmx-web-go/internal/services/bridge"
	"github.com/H3rby7/dmx-web-go/internal/services/config"
	"github.com/H3rby7/dmx-web-go/internal/services/fading"
	log "github.com/sirupsen/logrus"
)

// Stateful construct for chases
type ChaseService struct {
	chases        []models_chase.Chase
	fadingService *fading.FadingService
	bridgeService *bridge.BridgeService
}

func NewChaseService(configService *config.ConfigService, fadingService *fading.FadingService, bridgeService *bridge.BridgeService) *ChaseService {
	log.Debugf("Creating new ChaseService")
	chases := configService.GetChases()
	return &ChaseService{
		chases:        chases,
		fadingService: fadingService,
		bridgeService: bridgeService,
	}
}

/*
Reset a chase to its first element and start running it.
*/
func (svc *ChaseService) StartChaseFromTheTop(chaseName string) (ok bool) {
	ll := log.WithField("chase", chaseName)
	ok, chase := svc.findChaseByName(chaseName)
	if !ok {
		ll.Warnf("Could not find chase")
		return
	}
	ll.Debugf("Found chase")
	chase.RunFromStart(svc.renderDelegate, svc.bridgeDelegate)
	return
}

func (svc *ChaseService) findChaseByName(chaseName string) (ok bool, chase models_chase.Chase) {
	ok = false
	for _, c := range svc.chases {
		if c.Name == chaseName {
			ok = true
			chase = c
			return
		}
	}
	return
}

func (svc *ChaseService) renderDelegate(scene models_scene.Scene, fadeTimeMillis int64) {
	if svc.fadingService == nil {
		log.Warnf("No FadingService, skipping... ")
		return
	}
	svc.fadingService.FadeScene(scene, fadeTimeMillis)
}

func (svc *ChaseService) bridgeDelegate(active bool) {
	if svc.bridgeService == nil {
		log.Warnf("No DMX bridge, skipping... ")
		return
	}
	log.Debugf("Setting Bridge State... ")
	if active {
		svc.bridgeService.Activate()
	} else {
		svc.bridgeService.Deactivate()
	}
}
