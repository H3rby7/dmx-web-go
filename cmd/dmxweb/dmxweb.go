package main

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"

	apiv1 "github.com/H3rby7/dmx-web-go/internal/api/v1"
	dmxconn "github.com/H3rby7/dmx-web-go/internal/dmx"
	"github.com/H3rby7/dmx-web-go/internal/options"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func SetUpLogging() {
	log.SetLevel(options.GetDMXWebOptions().LogLevel)
	var formatter = nested.Formatter{
		// HideKeys:        true,
		CallerFirst:     true,
		FieldsOrder:     []string{"time", "component", "category"},
		TimestampFormat: time.RFC3339,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			filename := path.Base(f.File)
			fun := strings.Replace(f.Function, "github.com/H3rby7/dmx-web-go", "", 1)
			return fmt.Sprintf(" %s:%d::%s()", filename, f.Line, fun)
		},
	}
	log.SetFormatter(&formatter)
	log.SetReportCaller(true)
}

func main() {
	SetUpLogging()
	opts := options.GetDMXWebOptions()
	dmxOut := dmxconn.GetDMXConn()
	dmxOut.Connect(opts.SerialPort)

	router := gin.Default()
	apiv1.RegisterHandlers(router.Group("/api/v1"))
	webPort := opts.HttpPort
	log.Warnf("Serving at :%s", webPort)
	_ = router.Run(":" + webPort)
}
