package setup

import (
	models_services "github.com/H3rby7/dmx-web-go/internal/model/services"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/dmx-web-go/internal/services/bridge"
	"github.com/H3rby7/dmx-web-go/internal/services/chase"
	"github.com/H3rby7/dmx-web-go/internal/services/config"
	"github.com/H3rby7/dmx-web-go/internal/services/dmx"
	"github.com/H3rby7/dmx-web-go/internal/services/fading"
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

	services.DMXReaderService = dmx.NewDMXReaderService()
	services.DMXWriterService = dmx.NewDMXWriterService()
	services.FadingService = fading.NewFadingService(services.DMXWriterService)
	services.FadingService.Start()

	if opts.DmxBridge {
		if ok, objection := opts.CanBridge(); ok {
			services.BridgeService = bridge.NewBridgeService(services.DMXReaderService, services.FadingService)
			services.BridgeService.Activate()
			go services.BridgeService.BridgeDMX()
		} else {
			log.Infof("%s -> Skipping creation of 'BridgeService'", objection)
		}
	}

	services.ConfigService = config.NewConfigService()
	services.ChaseService = chase.NewChaseService(services.ConfigService, services.FadingService)
	services.TriggerService = trigger.NewTriggerService(services.ConfigService, services.ChaseService)
	return services
}
