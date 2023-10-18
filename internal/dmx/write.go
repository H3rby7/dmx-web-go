package dmxconn

import (
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

func GetWriter() *dmxusbpro.EnttecDMXUSBProController {
	return dmxConn.writer
}

func GetReader() *dmxusbpro.EnttecDMXUSBProController {
	return dmxConn.reader
}

func Initialize() {
	dmxConn.initWriter()
	dmxConn.initReader()
}

func (d *DMXConn) initWriter() {
	opts := options.GetAppOptions()
	port := opts.DmxWritePort
	baud := opts.DmxWriteBaudrate
	log.Infof("Opening DMX Serial for WRITING using port %s", port)
	config := &serial.Config{Name: port, Baud: baud}

	// Create a controller and connect to it
	d.writer = dmxusbpro.NewEnttecDMXUSBProController(config, true)
	if err := d.writer.Connect(); err != nil {
		log.Fatalf("Failed to connect DMX Controller for WRITING: %s", err)
	}
}

func (d *DMXConn) initReader() {
	opts := options.GetAppOptions()
	port := opts.DmxReadPort
	if port == "" {
		log.Warnf("No port specified to READ from DMX - skipping")
		return
	}
	baud := opts.DmxReadBaudrate
	log.Infof("Opening DMX Serial for READING using port %s", port)
	config := &serial.Config{Name: port, Baud: baud}

	// Create a controller and connect to it
	d.reader = dmxusbpro.NewEnttecDMXUSBProController(config, false)
	if err := d.reader.Connect(); err != nil {
		log.Fatalf("Failed to connect DMX Controller for READING: %s", err)
	}
}
