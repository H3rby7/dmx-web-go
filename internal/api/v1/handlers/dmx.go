package apiv1handlers

import (
	models "github.com/H3rby7/dmx-web-go/internal/api/v1/models"
	services "github.com/H3rby7/dmx-web-go/internal/api/v1/services"
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
	data := models.Scene{}
	err := c.BindJSON(&data)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(400)
		return
	}
	services.SetScene(data, 0)
	c.String(200, "OK")
}

// Apply DMX values with fadetime
func patchDmxFade(c *gin.Context) {
	data := models.SceneWithFade{}
	err := c.BindJSON(&data)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(400)
		return
	}
	services.SetScene(data.Scene, data.FadeTimeMillis)
	c.String(200, "OK")
}

// Clear all DMX values immediately
func putDmxClear(c *gin.Context) {
	services.ClearAll()
	c.String(200, "OK")
}

// Extract a Scene strcut from the gin request
func sceneFromJson(c *gin.Context) (s models.Scene, err error) {
	s = models.Scene{}
	err = c.BindJSON(&s)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(400)
	}
	return
}
