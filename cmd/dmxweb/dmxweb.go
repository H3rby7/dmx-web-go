package main

import (
	"flag"

	apiv1 "github.com/H3rby7/dmx-web-go/internal/api/v1"
	"github.com/H3rby7/dmx-web-go/internal/setup"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	in := flag.String("dmx-in", "", "Input interface (e.g. COM4 OR /dev/tty.usbserial)")
	out := flag.String("dmx-out", "", "Output interface (e.g. COM4 OR /dev/tty.usbserial)")
	baud := flag.Int("baud", 57600, "Baudrate for the device")
	webPort := *flag.String("port", "8080", "Port for the HTTP server")
	loglevel := *flag.String("loglevel", "info", "Baudrate for the device")
	flag.Parse()

	setup.SetUpLogging(loglevel)

	router := gin.Default()
	apiv1.RegisterHandlers(router.Group("/api/v1"))
	log.Warnf("Serving at :%s", webPort)
	_ = router.Run(":" + webPort)
}
