package apiv1handlers

import (
	"github.com/H3rby7/dmx-web-go/internal/services/bridge"
	"github.com/gin-gonic/gin"
)

// Register Handlers for V1 API
func RegisterBridgeHandlers(g *gin.RouterGroup, svc *bridge.BridgeService) {
	g.PUT("bridge/activate", func(ctx *gin.Context) { svc.Activate() })
	g.PUT("bridge/deactivate", func(ctx *gin.Context) { svc.Deactivate() })
}
