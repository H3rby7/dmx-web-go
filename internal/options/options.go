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
	DmxClearOnQuit   bool
}

// Local instance holding our settings
var optionsInstance = AppOptions{}

// Define application's flags, parse them and read them into our options struct.
func InitAppOptions() {
	var err error
	httpPort := flag.String("http-port", "8080", "HTTP Server port")
	dmxReadBaudrate := flag.Int("dmx-read-baud", 57600, "Baudrate for the reading device")
	dmxReadPort := flag.String("dmx-read-port", "", "Input interface (e.g. COM4 OR /dev/tty.usbserial)")
	dmxWriteBaudrate := flag.Int("dmx-write-baud", 57600, "Baudrate for the writing device")
	dmxWritePort := flag.String("dmx-write-port", "", "Output interface (e.g. COM4 OR /dev/tty.usbserial)")
	dmxClearOnQuit := flag.Bool("dmx-clear-on-quit", true, "Whether or not to send '0's out for all DMX channels upon exit.")
	logLevel := flag.String("log-level", "info", "Granularity of log output, see logrus.ParseLevel")
	flag.Parse()
	optionsInstance.HttpPort = *httpPort
	optionsInstance.DmxReadBaudrate = *dmxReadBaudrate
	optionsInstance.DmxReadPort = *dmxReadPort
	optionsInstance.DmxWriteBaudrate = *dmxWriteBaudrate
	optionsInstance.DmxWritePort = *dmxWritePort
	optionsInstance.DmxClearOnQuit = *dmxClearOnQuit
	optionsInstance.LogLevel, err = log.ParseLevel(*logLevel)
	if err != nil {
		optionsInstance.LogLevel = log.InfoLevel
	}
}

// Get the options
func GetAppOptions() AppOptions {
	return optionsInstance
}
