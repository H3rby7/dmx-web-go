package apiv1

import (
	apiv1handlers "github.com/H3rby7/dmx-web-go/internal/api/v1/handlers"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/dmx-web-go/internal/services/trigger"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Register Handlers for V1 API
func RegisterHandlers(g *gin.RouterGroup, triggerSvc *trigger.TriggerService) {
	opts := options.GetAppOptions()

	if ok, objection := opts.CanBridge(); ok {
		apiv1handlers.RegisterBridgeHandlers(g)
	} else {
		log.Warnf("%s -> Skipping registration of 'bridge' API", objection)
	}

	if ok, objection := opts.CanWriteDMX(); ok {
		apiv1handlers.RegisterDMXHandlers(g)
	} else {
		log.Warnf("%s -> Skipping registration of 'DMX' API", objection)
	}

	apiv1handlers.RegisterTriggerHandlers(g, triggerSvc)

}
