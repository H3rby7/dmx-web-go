// Package chase hosts the ChaseService and provides means to work with Chases
package chase

import (
	models_chase "github.com/H3rby7/dmx-web-go/internal/model/chase"
	models_scene "github.com/H3rby7/dmx-web-go/internal/model/scene"
	"github.com/H3rby7/dmx-web-go/internal/services/bridge"
	"github.com/H3rby7/dmx-web-go/internal/services/config"
	"github.com/H3rby7/dmx-web-go/internal/services/fading"
	log "github.com/sirupsen/logrus"
)

// ChaseService handles chases and actions on chases
type ChaseService struct {
	chases        []models_chase.Chase
	fadingService *fading.FadingService
	bridgeService *bridge.BridgeService
}

// NewChaseService creates a new [ChaseService] instance
//
// Also loads the chases from the [ConfigService]
func NewChaseService(configService *config.ConfigService, fadingService *fading.FadingService, bridgeService *bridge.BridgeService) *ChaseService {
	log.Debugf("Creating new ChaseService")
	chases := configService.GetChases()
	return &ChaseService{
		chases:        chases,
		fadingService: fadingService,
		bridgeService: bridgeService,
	}
}

// StartChaseFromTheTop resets a chase to its first element and start running it.
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

// findChaseByName looks for a chase matching the name
//
// Returns a reference to the given chase, if found.
// That way the same instance of the chase is returned on censecutive calls with identical chaseName
func (svc *ChaseService) findChaseByName(chaseName string) (ok bool, chase *models_chase.Chase) {
	ok = false
	for i := 0; i < len(svc.chases); i++ {
		if svc.chases[i].Name == chaseName {
			ok = true
			chase = &svc.chases[i]
			return
		}
	}
	return
}

// renderDelegate implements [models_chase.SceneRenderFunc]
func (svc *ChaseService) renderDelegate(scene models_scene.Scene, fadeTimeMillis int64) {
	svc.fadingService.FadeScene(scene, fadeTimeMillis)
}

// bridgeDelegate implements [models_chase.ChangeBridgeStateFunc]
func (svc *ChaseService) bridgeDelegate(active bool, fadeDurationMillis int64) {
	log.Debugf("Setting Bridge State... ")
	if active {
		svc.bridgeService.Activate(fadeDurationMillis)
	} else {
		svc.bridgeService.Deactivate(fadeDurationMillis)
	}
}
