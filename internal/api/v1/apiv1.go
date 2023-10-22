package apiv1

import (
	apiv1handlers "github.com/H3rby7/dmx-web-go/internal/api/v1/handlers"
	"github.com/gin-gonic/gin"
)

// Register Handlers for V1 API
func RegisterHandlers(g *gin.RouterGroup) {
	apiv1handlers.RegisterBridgeHandlers(g)
	apiv1handlers.RegisterDMXHandlers(g)
}
