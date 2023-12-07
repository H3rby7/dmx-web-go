package setup

import (
	models_services "github.com/H3rby7/dmx-web-go/internal/model/services"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/dmx-web-go/internal/services/bridge"
	"github.com/H3rby7/dmx-web-go/internal/services/chase"
	"github.com/H3rby7/dmx-web-go/internal/services/config"
	"github.com/H3rby7/dmx-web-go/internal/services/trigger"
	log "github.com/sirupsen/logrus"
)

// InitServices creates and initializes all services needed by the application
//
// Returns a struct of type [ApplicationServices] holding all service references
func InitServices() *models_services.ApplicationServices {
	log.Infof("Initializing Application Services... ")
	services := &models_services.ApplicationServices{}
	opts := options.GetAppOptions()

	if ok, objection := opts.CanBridge(); ok {
		services.BridgeService = bridge.NewBridgeService()
	} else {
		log.Infof("%s -> Skipping creation of 'BridgeService'", objection)
	}

	services.ConfigService = config.NewConfigService()
	services.ChaseService = chase.NewChaseService(services.ConfigService)
	services.TriggerService = trigger.NewTriggerService(services.ConfigService, services.ChaseService)
	return services
}
