package apiv1handlers

import (
	dtos "github.com/H3rby7/dmx-web-go/internal/api/v1/dtos"
	apiv1mappers "github.com/H3rby7/dmx-web-go/internal/api/v1/mappers"
	services "github.com/H3rby7/dmx-web-go/internal/services"
	"github.com/gin-gonic/gin"
)

// Register Handlers for V1 API
func RegisterDMXHandlers(g *gin.RouterGroup) {
	g.PATCH("dmx", patchDmx)
	g.PATCH("dmx/fade", patchDmxFade)
	g.PUT("dmx/clear", putDmxClear)
}

// Apply DMX values immediately
func patchDmx(c *gin.Context) {
	data := dtos.Scene{}
	err := c.BindJSON(&data)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(400)
		return
	}
	services.SetScene(apiv1mappers.MapScene(data), 0)
	c.String(200, "OK")
}

// Apply DMX values with fadetime
func patchDmxFade(c *gin.Context) {
	data := dtos.SceneWithFade{}
	err := c.BindJSON(&data)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(400)
		return
	}
	services.SetScene(apiv1mappers.MapScene(data.Scene), data.FadeTimeMillis)
	c.String(200, "OK")
}

// Clear all DMX values immediately
func putDmxClear(c *gin.Context) {
	services.ClearAll()
	c.String(200, "OK")
}

// Extract a Scene strcut from the gin request
func sceneFromJson(c *gin.Context) (s dtos.Scene, err error) {
	s = dtos.Scene{}
	err = c.BindJSON(&s)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(400)
	}
	return
}
