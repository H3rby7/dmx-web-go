package main

import (
	apiv1 "github.com/H3rby7/dmx-web-go/internal/api/v1"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/dmx-web-go/internal/setup"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	opts := options.InitAppOptions()
	setup.SetUpLogging(opts.LogLevel)

	router := gin.Default()
	apiv1.RegisterHandlers(router.Group("/api/v1"))
	log.Warnf("Serving at :%s", opts.HttpPort)
	_ = router.Run(":" + opts.HttpPort)
}
