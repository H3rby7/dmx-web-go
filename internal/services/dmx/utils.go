package dmx

import (
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/sirupsen/logrus"
)

/*
Returns the extracted log-level in DMX USB log verbosity format

See https://github.com/H3rby7/usbdmx-golang/blob/main/usbdmxgolang.go > SetLogVerbosity(unit8)
*/
func logVerbosity() uint8 {
	level := options.GetAppOptions().LogLevel
	switch level {
	case logrus.TraceLevel:
		return 2
	case logrus.DebugLevel:
		return 1
	default:
		return 0
	}
}
