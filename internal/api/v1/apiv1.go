package apiv1

import (
	apiv1handlers "github.com/H3rby7/dmx-web-go/internal/api/v1/handlers"
	models_services "github.com/H3rby7/dmx-web-go/internal/model/services"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Register Handlers for V1 API
func RegisterHandlers(g *gin.RouterGroup, services *models_services.ApplicationServices) {
	ll := log.WithField("API", "V1")
	ll.Debugf("Registering handlers... ")
	opts := options.GetAppOptions()

	if ok, objection := opts.CanBridge(); ok {
		ll.Infof("Registering handlers for 'BRIDGE'")
		apiv1handlers.RegisterBridgeHandlers(g, services.BridgeService)
	} else {
		ll.Infof("%s -> Skipping registration of handlers for 'BRIDGE'", objection)
	}

	if ok, objection := opts.CanWriteDMX(); ok {
		ll.Infof("Registering handlers for 'DMX'")
		apiv1handlers.RegisterDMXHandlers(g, services.FadingService)
	} else {
		ll.Infof("%s -> Skipping registration of handlers for 'DMX'", objection)
	}

	ll.Infof("Registering handlers for 'TRIGGER'")
	apiv1handlers.RegisterTriggerHandlers(g, services.TriggerService)

	ll.Infof("Done registering handlers for API V1")
}
