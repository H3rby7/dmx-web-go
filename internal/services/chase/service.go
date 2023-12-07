package chase

import (
	models_chase "github.com/H3rby7/dmx-web-go/internal/model/chase"
	models_scene "github.com/H3rby7/dmx-web-go/internal/model/scene"
	"github.com/H3rby7/dmx-web-go/internal/services"
	"github.com/H3rby7/dmx-web-go/internal/services/config"
	log "github.com/sirupsen/logrus"
)

// Stateful construct for chases
type ChaseService struct {
	chases []models_chase.Chase
}

func NewChaseService(configService *config.ConfigService) *ChaseService {
	log.Debugf("Creating new ChaseService")
	chases := configService.GetChases()
	return &ChaseService{
		chases: chases,
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
	chase.RunFromStart(func(scene models_scene.Scene, fadeTimeMillis int64) {
		ll.Debugf("Rendering... ")
		services.SetScene(scene, fadeTimeMillis)
	})
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
