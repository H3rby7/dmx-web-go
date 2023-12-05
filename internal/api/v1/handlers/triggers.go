package apiv1handlers

import (
	models "github.com/H3rby7/dmx-web-go/internal/api/v1/models"
	services "github.com/H3rby7/dmx-web-go/internal/api/v1/services"
	"github.com/gin-gonic/gin"
)

// Register Handlers for V1 API
func RegisterTriggerHandlers(g *gin.RouterGroup, svc *services.TriggerService) {
	g.POST("trigger", createTriggerHandler(svc))
}

// Handle triggers from ?
func createTriggerHandler(svc *services.TriggerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tr := models.TriggerSource{}
		err := c.BindJSON(&tr)
		if err != nil {
			c.Error(err)
			c.AbortWithStatus(400)
			return
		}
		svc.Handle(tr.Source)
		c.String(200, "OK")
	}
}
