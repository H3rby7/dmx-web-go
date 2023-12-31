// Package setup contains configurations for logging and server as well as service creation
package setup

import (
	models_services "github.com/H3rby7/dmx-web-go/internal/model/services"
	"github.com/H3rby7/dmx-web-go/internal/services/bridge"
	"github.com/H3rby7/dmx-web-go/internal/services/chase"
	"github.com/H3rby7/dmx-web-go/internal/services/config"
	"github.com/H3rby7/dmx-web-go/internal/services/event"
	"github.com/H3rby7/dmx-web-go/internal/services/fading"
	"github.com/H3rby7/dmx-web-go/internal/services/reader"
	"github.com/H3rby7/dmx-web-go/internal/services/trigger"
	log "github.com/sirupsen/logrus"
)

// InitServices creates and initializes all services needed by the application
//
// Returns a struct of type [ApplicationServices] holding all service references
func InitServices() *models_services.ApplicationServices {
	log.Infof("Initializing Application Services... ")
	services := &models_services.ApplicationServices{}

	services.DMXReaderService = reader.NewDMXReaderService()
	services.FadingService = fading.NewFadingService()
	services.BridgeService = bridge.NewBridgeService(services.DMXReaderService, services.FadingService)

	services.ConfigService = config.NewConfigService()
	services.ChaseService = chase.NewChaseService(services.ConfigService, services.FadingService, services.BridgeService)
	services.EventService = event.NewEventService(services.ConfigService, services.ChaseService)
	services.TriggerService = trigger.NewTriggerService(services.ConfigService, services.EventService)
	return services
}
