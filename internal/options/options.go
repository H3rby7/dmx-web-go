package options

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

type AppOptions struct {
	HttpPort         string
	Static           string
	LogLevel         log.Level
	DmxLogLevel      uint8
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
	static := flag.String("static", "", "Relative path for static file serving. Leave blank for no serving.")
	dmxChannels := flag.Int("dmx-channels", 512, "DMX channel count. Lower count saves some energy as less needs to be sent")
	dmxReadBaudrate := flag.Int("dmx-read-baud", 57600, "Baudrate for the reading device")
	dmxReadPort := flag.String("dmx-read-port", "", "Input interface (e.g. COM4 OR /dev/tty.usbserial)")
	dmxWriteBaudrate := flag.Int("dmx-write-baud", 57600, "Baudrate for the writing device")
	dmxWritePort := flag.String("dmx-write-port", "", "Output interface (e.g. COM4 OR /dev/tty.usbserial)")
	dmxClearOnQuit := flag.Bool("dmx-clear-on-quit", true, "Whether or not to send '0's out for all DMX channels upon exit.")
	dmxBridge := flag.Bool("dmx-bridge", false, "Whether or not to send the read input into write upon receiving.")
	dmxLogLevel := flag.Uint("dmx-log-level", 0, "Granularity of DMX logs [0,1,2]")
	logLevel := flag.String("log-level", "info", "Granularity of log output, see logrus.ParseLevel")
	flag.Parse()
	optionsInstance.HttpPort = *httpPort
	optionsInstance.Static = *static
	optionsInstance.DmxChannelCount = *dmxChannels
	optionsInstance.DmxReadBaudrate = *dmxReadBaudrate
	optionsInstance.DmxReadPort = *dmxReadPort
	optionsInstance.DmxWriteBaudrate = *dmxWriteBaudrate
	optionsInstance.DmxWritePort = *dmxWritePort
	optionsInstance.DmxClearOnQuit = *dmxClearOnQuit
	optionsInstance.DmxBridge = *dmxBridge
	optionsInstance.DmxLogLevel = uint8(*dmxLogLevel)
	optionsInstance.LogLevel, err = log.ParseLevel(*logLevel)
	if err != nil {
		optionsInstance.LogLevel = log.InfoLevel
	}
}

// Get the options
func GetAppOptions() AppOptions {
	return optionsInstance
}

func (opts *AppOptions) CanWriteDMX() (ok bool, objection string) {
	if opts.DmxWritePort == "" {
		return false, "No DMX writer specified"
	}
	return true, ""
}

func (opts *AppOptions) CanReadDMX() (ok bool, objection string) {
	if opts.DmxReadPort == "" {
		return false, "No DMX reader specified"
	}
	return true, ""
}

func (opts *AppOptions) CanBridge() (ok bool, objection string) {
	if !opts.DmxBridge {
		return false, "DMX-Bridge flag is FALSE"
	}
	if ok, writeObj := opts.CanWriteDMX(); !ok {
		return false, writeObj
	}
	if ok, readObj := opts.CanReadDMX(); !ok {
		return false, readObj
	}
	return true, ""
}
