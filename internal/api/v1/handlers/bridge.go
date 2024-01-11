// Package apiv1handlers defines all handlers used by the API in version 1
package apiv1handlers

import (
	models_fader "github.com/H3rby7/dmx-web-go/internal/model/fader"
	"github.com/H3rby7/dmx-web-go/internal/services/bridge"
	"github.com/gin-gonic/gin"
)

// RegisterBridgeHandlers registers the bridge handlers for V1 API
func RegisterBridgeHandlers(g *gin.RouterGroup, svc *bridge.BridgeService) {
	g.PUT("bridge/activate", func(ctx *gin.Context) { svc.Activate(models_fader.FADE_IMMEDIATELY) })
	g.PUT("bridge/deactivate", func(ctx *gin.Context) { svc.Deactivate(models_fader.FADE_IMMEDIATELY) })
}
