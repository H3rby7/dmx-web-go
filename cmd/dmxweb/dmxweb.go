package main

import (
	"fmt"
	"net/http"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/H3rby7/dmx-web-go/internal/options"
	nested "github.com/antonfisher/nested-logrus-formatter"
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
	port := options.GetDMXWebOptions().HttpPort
	log.Warn("Serving at :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
