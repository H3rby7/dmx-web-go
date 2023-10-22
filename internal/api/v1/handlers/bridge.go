package apiv1handlers

import (
	dmxconn "github.com/H3rby7/dmx-web-go/internal/dmx"
	"github.com/gin-gonic/gin"
)

// Register Handlers for V1 API
func RegisterBridgeHandlers(g *gin.RouterGroup) {
	g.PUT("bridge/activate", bridgeActivateHandler)
	g.PUT("bridge/deactivate", bridgeDeactivateHandler)
}

func bridgeActivateHandler(c *gin.Context) {
	dmxconn.GetBridge().Activate()
}

func bridgeDeactivateHandler(c *gin.Context) {
	dmxconn.GetBridge().Deactivate()
}
