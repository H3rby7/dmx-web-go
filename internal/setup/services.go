// Package setup contains configurations for logging and server as well as service creation
package setup

import (
	models_services "github.com/H3rby7/dmx-web-go/internal/model/services"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/dmx-web-go/internal/services/bridge"
	"github.com/H3rby7/dmx-web-go/internal/services/chase"
	"github.com/H3rby7/dmx-web-go/internal/services/config"
	"github.com/H3rby7/dmx-web-go/internal/services/dmxmock"
	"github.com/H3rby7/dmx-web-go/internal/services/enttec/dmxusbpro"
	"github.com/H3rby7/dmx-web-go/internal/services/event"
	"github.com/H3rby7/dmx-web-go/internal/services/printer"
	"github.com/H3rby7/dmx-web-go/internal/services/trigger"
	log "github.com/sirupsen/logrus"
)

// InitServices creates and initializes all services needed by the application
//
// Returns a struct of type [ApplicationServices] holding all service references
func InitServices() *models_services.ApplicationServices {
	opts := options.GetAppOptions()
	log.Infof("Initializing Application Services... ")
	services := &models_services.ApplicationServices{}

	if opts.ReadUsesMock() {
		services.DMXReaderService = dmxmock.NewMockedDMXReaderService()
	} else {
		services.DMXReaderService = dmxusbpro.NewDMXReaderService()
	}
	if opts.WriteUsesMock() {
		services.FadingService = dmxmock.NewMockedFadingService()
	} else {
		services.FadingService = dmxusbpro.NewFadingService()
	}
	services.BridgeService = bridge.NewBridgeService(services.DMXReaderService, services.FadingService)

	if willBridge, _ := opts.CanBridge(); !willBridge {
		if canRead, _ := opts.CanReadDMX(); canRead {
			log.Infof("Creating DMX Logger to utilize the READability.")
			pSvc := printer.NewDMXLoggerService(services.DMXReaderService)
			go pSvc.PrintDMX()
		}
	}

	services.ConfigService = config.NewConfigService()
	services.ChaseService = chase.NewChaseService(services.ConfigService, services.FadingService, services.BridgeService)
	services.EventService = event.NewEventService(services.ConfigService, services.ChaseService)
	services.TriggerService = trigger.NewTriggerService(services.ConfigService, services.EventService)
	return services
}
