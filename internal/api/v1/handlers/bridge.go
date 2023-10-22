package apiv1handlers

import (
	dmxconn "github.com/H3rby7/dmx-web-go/internal/dmx"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Register Handlers for V1 API
func RegisterBridgeHandlers(g *gin.RouterGroup) {
	if !options.GetAppOptions().DmxBridge {
		logrus.Warnf("DMX-Bridge flag is false -> Skipping registration of bridge API")
		return
	}
	g.PUT("bridge/activate", bridgeActivateHandler)
	g.PUT("bridge/deactivate", bridgeDeactivateHandler)
}

func bridgeActivateHandler(c *gin.Context) {
	dmxconn.GetBridge().Activate()
}

func bridgeDeactivateHandler(c *gin.Context) {
	dmxconn.GetBridge().Deactivate()
}
