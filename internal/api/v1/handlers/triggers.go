package apiv1handlers

import (
	dtos "github.com/H3rby7/dmx-web-go/internal/api/v1/dtos"
	"github.com/H3rby7/dmx-web-go/internal/services/trigger"
	"github.com/gin-gonic/gin"
)

// Register Handlers for V1 API
func RegisterTriggerHandlers(g *gin.RouterGroup, svc *trigger.TriggerService) {
	g.POST("trigger", createTriggerHandler(svc))
}

// createTriggerHandler returns a HandlerFunction as needed by GIN
//
// The function parses the request and handles it using the [TriggerService]
func createTriggerHandler(svc *trigger.TriggerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tr := dtos.TriggerSource{}
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
