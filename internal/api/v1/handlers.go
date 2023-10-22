package apiv1

import (
	dmxconn "github.com/H3rby7/dmx-web-go/internal/dmx"
	"github.com/gin-gonic/gin"
)

// Register Handlers for V1 API
func RegisterHandlers(g *gin.RouterGroup) {
	g.POST("dmx", dmxHandler)
	g.POST("bridge/activate", bridgeActivateHandler)
	g.POST("bridge/deactivate", bridgeDeactivateHandler)
}

// Handle incoming dmx channel
func dmxHandler(c *gin.Context) {
	data := MultipleDMXValueForChannel{}
	err := c.BindJSON(&data)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(400)
		return
	}
	dmxSetMulti(data)
	c.String(200, "OK")
}

func bridgeActivateHandler(c *gin.Context) {
	dmxconn.GetBridge().Activate()
}

func bridgeDeactivateHandler(c *gin.Context) {
	dmxconn.GetBridge().Deactivate()
}
