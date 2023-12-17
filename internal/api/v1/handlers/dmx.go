// Package apiv1handlers defines all handlers used by the API in version 1
package apiv1handlers

import (
	dtos "github.com/H3rby7/dmx-web-go/internal/api/v1/dtos"
	apiv1mappers "github.com/H3rby7/dmx-web-go/internal/api/v1/mappers"
	models_fader "github.com/H3rby7/dmx-web-go/internal/model/fader"
	"github.com/H3rby7/dmx-web-go/internal/services/fading"
	"github.com/gin-gonic/gin"
)

// RegisterDMXHandlers registers the DMX handlers for V1 API
func RegisterDMXHandlers(g *gin.RouterGroup, svc *fading.FadingService) {
	g.PATCH("dmx", createPatchDmxHandler(svc))
	g.PATCH("dmx/fade", createPatchDmxFadeHandler(svc))
	g.PUT("dmx/clear", createPutDmxClearHandler(svc))
}

// createPatchDmxHandler returns a HandlerFunction as needed by GIN
//
// The function parses the request body and applies the scene immediately
func createPatchDmxHandler(svc *fading.FadingService) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := dtos.Scene{}
		err := c.BindJSON(&data)
		if err != nil {
			c.Error(err)
			c.AbortWithStatus(400)
			return
		}
		svc.FadeScene(apiv1mappers.MapScene(data), models_fader.FADE_IMMEDIATELY)
		c.String(200, "OK")
	}
}

// createPatchDmxFadeHandler returns a HandlerFunction as needed by GIN
//
// The function parses the request and and applies the scene using the given fade time
func createPatchDmxFadeHandler(svc *fading.FadingService) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := dtos.SceneWithFade{}
		err := c.BindJSON(&data)
		if err != nil {
			c.Error(err)
			c.AbortWithStatus(400)
			return
		}
		svc.FadeScene(apiv1mappers.MapScene(data.Scene), data.FadeTimeMillis)
		c.String(200, "OK")
	}
}

// Clear all DMX values immediately
func createPutDmxClearHandler(svc *fading.FadingService) gin.HandlerFunc {
	return func(c *gin.Context) {
		svc.ClearAll()
		c.String(200, "OK")
	}
}
