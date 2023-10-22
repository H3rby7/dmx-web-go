package options

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

type AppOptions struct {
	HttpPort         string
	LogLevel         log.Level
	DmxChannelCount  int
	DmxWritePort     string
	DmxWriteBaudrate int
	DmxReadPort      string
	DmxReadBaudrate  int
	DmxClearOnQuit   bool
	DmxBridge        bool
}

// Local instance holding our settings
var optionsInstance = AppOptions{}

// Define application's flags, parse them and read them into our options struct.
func InitAppOptions() {
	var err error
	httpPort := flag.String("http-port", "8080", "HTTP Server port")
	dmxChannels := flag.Int("dmx-channels", 512, "DMX channel count. Lower count saves some energy as less needs to be sent")
	dmxReadBaudrate := flag.Int("dmx-read-baud", 57600, "Baudrate for the reading device")
	dmxReadPort := flag.String("dmx-read-port", "", "Input interface (e.g. COM4 OR /dev/tty.usbserial)")
	dmxWriteBaudrate := flag.Int("dmx-write-baud", 57600, "Baudrate for the writing device")
	dmxWritePort := flag.String("dmx-write-port", "", "Output interface (e.g. COM4 OR /dev/tty.usbserial)")
	dmxClearOnQuit := flag.Bool("dmx-clear-on-quit", true, "Whether or not to send '0's out for all DMX channels upon exit.")
	dmxBridge := flag.Bool("dmx-bridge", false, "Whether or not to send the read input into write upon receiving.")
	logLevel := flag.String("log-level", "info", "Granularity of log output, see logrus.ParseLevel")
	flag.Parse()
	optionsInstance.HttpPort = *httpPort
	optionsInstance.DmxChannelCount = *dmxChannels
	optionsInstance.DmxReadBaudrate = *dmxReadBaudrate
	optionsInstance.DmxReadPort = *dmxReadPort
	optionsInstance.DmxWriteBaudrate = *dmxWriteBaudrate
	optionsInstance.DmxWritePort = *dmxWritePort
	optionsInstance.DmxClearOnQuit = *dmxClearOnQuit
	optionsInstance.DmxBridge = *dmxBridge
	optionsInstance.LogLevel, err = log.ParseLevel(*logLevel)
	if err != nil {
		optionsInstance.LogLevel = log.InfoLevel
	}
}

// Get the options
func GetAppOptions() AppOptions {
	return optionsInstance
}

func (opts *AppOptions) HasDMXWriter() bool {
	return opts.DmxWritePort != ""
}

func (opts *AppOptions) HasDMXReader() bool {
	return opts.DmxReadPort != ""
}
