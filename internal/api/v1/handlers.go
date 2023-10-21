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
	g.POST("testsolo", testsoloHandler)
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
	dmx := dmxconn.GetWriter()
	for _, entry := range data.List {
		if err = dmx.Stage(entry.Channel, entry.Value); err != nil {
			c.Error(err)
		}
	}
	if len(c.Errors) != 0 {
		c.AbortWithStatus(400)
		return
	}
	err = dmx.Commit()
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(500)
		return
	}
	c.String(200, "OK")
}

func bridgeActivateHandler(c *gin.Context) {
	dmxconn.GetBridge().Activate()
}

func bridgeDeactivateHandler(c *gin.Context) {
	dmxconn.GetBridge().Deactivate()
}

func testsoloHandler(c *gin.Context) {
	bridgeDeactivateHandler(c)
	w := dmxconn.GetWriter()
	w.ClearStage()
	data := MultipleDMXValueForChannel{
		List: []DMXValueForChannel{
			{15, 150},
		},
	}
	dmx := dmxconn.GetWriter()
	for _, entry := range data.List {
		if err := dmx.Stage(entry.Channel, entry.Value); err != nil {
			c.Error(err)
		}
	}
	if len(c.Errors) != 0 {
		c.AbortWithStatus(400)
		return
	}
	err := dmx.Commit()
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(500)
		return
	}
	c.String(200, "OK")
}
