package options

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

type AppOptions struct {
	HttpPort         string
	LogLevel         log.Level
	DmxWritePort     string
	DmxWriteBaudrate int
	DmxReadPort      string
	DmxReadBaudrate  int
}

// Local instance holding our settings
var optionsInstance = AppOptions{}

func InitAppOptions() AppOptions {
	var err error
	optionsInstance.HttpPort = *flag.String("http-port", "8080", "HTTP Server port")
	optionsInstance.DmxReadBaudrate = *flag.Int("dmx-read-baud", 57600, "Baudrate for the reading device")
	optionsInstance.DmxReadPort = *flag.String("dmx-read-port", "", "Input interface (e.g. COM4 OR /dev/tty.usbserial)")
	optionsInstance.DmxWriteBaudrate = *flag.Int("dmx-write-baud", 57600, "Baudrate for the writing device")
	optionsInstance.DmxWritePort = *flag.String("dmx-write-port", "", "Output interface (e.g. COM4 OR /dev/tty.usbserial)")
	optionsInstance.LogLevel, err = log.ParseLevel(*flag.String("log-level", "info", "Granularity of log output, see logrus.ParseLevel"))
	if err != nil {
		optionsInstance.LogLevel = log.InfoLevel
	}
	flag.Parse()
	return optionsInstance
}

/*
Get Options
*/
func GetAppOptions() AppOptions {
	return optionsInstance
}
