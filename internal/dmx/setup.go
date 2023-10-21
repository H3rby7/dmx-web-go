package dmxconn

import (
	dmxbridge "github.com/H3rby7/dmx-web-go/internal/dmx/bridge"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/usbdmx-golang/controller/enttec/dmxusbpro"
	log "github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

type DMXConn struct {
	writer *dmxusbpro.EnttecDMXUSBProController
	reader *dmxusbpro.EnttecDMXUSBProController
}

var dmxConn = DMXConn{}
var bridge *dmxbridge.DMXBridge

func GetWriter() *dmxusbpro.EnttecDMXUSBProController {
	return dmxConn.writer
}

func GetReader() *dmxusbpro.EnttecDMXUSBProController {
	return dmxConn.reader
}

func GetBridge() *dmxbridge.DMXBridge {
	return bridge
}

func Initialize() {
	opts := options.GetAppOptions()
	dmxConn.initWriter()
	if opts.DmxReadPort != "" {
		dmxConn.initReader()
		GetReader().SetLogVerbosity(1)
		GetWriter().SetLogVerbosity(1)
		bridge = dmxbridge.NewDMXBridge(GetReader(), GetWriter())
		bridge.Activate()
		go bridge.BridgeDMX()
	}
}

func (d *DMXConn) initWriter() {
	opts := options.GetAppOptions()
	channels := opts.DmxChannelCount
	port := opts.DmxWritePort
	baud := opts.DmxWriteBaudrate
	log.Infof("Opening DMX Serial for WRITING using port %s", port)
	config := &serial.Config{Name: port, Baud: baud}

	// Create a controller and connect to it
	d.writer = dmxusbpro.NewEnttecDMXUSBProController(config, channels, true)
	if err := d.writer.Connect(); err != nil {
		log.Fatalf("Failed to connect DMX Controller for WRITING: %s", err)
	}
}

func (d *DMXConn) initReader() {
	opts := options.GetAppOptions()
	channels := opts.DmxChannelCount
	port := opts.DmxReadPort
	if port == "" {
		log.Warnf("No port specified to READ from DMX - skipping")
		return
	}
	baud := opts.DmxReadBaudrate
	log.Infof("Opening DMX Serial for READING using port %s", port)
	config := &serial.Config{Name: port, Baud: baud}

	// Create a controller and connect to it
	d.reader = dmxusbpro.NewEnttecDMXUSBProController(config, channels, false)
	if err := d.reader.Connect(); err != nil {
		log.Fatalf("Failed to connect DMX Controller for READING: %s", err)
	}
}
