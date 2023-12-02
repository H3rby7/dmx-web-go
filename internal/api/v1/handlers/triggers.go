package apiv1handlers

import (
	models "github.com/H3rby7/dmx-web-go/internal/api/v1/models"
	services "github.com/H3rby7/dmx-web-go/internal/api/v1/services"
	"github.com/gin-gonic/gin"
)

// Register Handlers for V1 API
func RegisterTriggerHandlers(g *gin.RouterGroup) {
	g.POST("trigger", handleTrigger)
}

// Handle triggers from ?
func handleTrigger(c *gin.Context) {
	tr := models.TriggerSource{}
	err := c.BindJSON(&tr)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(400)
		return
	}
	services.PrintHello()
	c.String(200, "OK")
}
