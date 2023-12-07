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
	log.Debugf("Registering handlers for API V1... ")
	opts := options.GetAppOptions()

	if ok, objection := opts.CanBridge(); ok {
		apiv1handlers.RegisterBridgeHandlers(g, services.BridgeService)
	} else {
		log.Warnf("%s -> Skipping registration of 'bridge' API", objection)
	}

	if ok, objection := opts.CanWriteDMX(); ok {
		apiv1handlers.RegisterDMXHandlers(g, services.FadingService)
	} else {
		log.Warnf("%s -> Skipping registration of 'DMX' API", objection)
	}

	apiv1handlers.RegisterTriggerHandlers(g, services.TriggerService)
	log.Infof("Done registering handlers for API V1")
}
