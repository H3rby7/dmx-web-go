// Package mockrest defines Handlers for DMX Mocks controlled via REST API
package mockrest

import (
	dtos "github.com/H3rby7/dmx-web-go/internal/api/v1/dtos"
	mock "github.com/H3rby7/dmx-web-go/internal/mock/services"
	models_services "github.com/H3rby7/dmx-web-go/internal/model/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// RegisterDMXHandlers registers the DMX handlers for V1 API
func RegisterMockHandlers(g *gin.RouterGroup, services *models_services.ApplicationServices) {
	if mock, ok := services.DMXReaderService.(*mock.MockedDMXReaderService); !ok {
		log.Fatalf("Expected DMXReaderService to be a mock!")
	} else {
		g.PATCH("read", createReadMockDmxHandler(*mock))
	}
}

// createReadMockDmxHandler returns a HandlerFunction as needed by GIN
//
// The function parses the request body and passes it as read input
func createReadMockDmxHandler(svc mock.MockedDMXReaderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := dtos.Scene{}
		err := c.BindJSON(&data)
		if err != nil {
			c.Error(err)
			c.AbortWithStatus(400)
			return
		}
		changeSet := make(map[int]byte)
		for _, el := range data.List {
			changeSet[int(el.Channel)] = el.Value
		}
		svc.ReadChangeset(changeSet)
		c.String(200, "OK")
	}
}
