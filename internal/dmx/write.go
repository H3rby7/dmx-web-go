package dmxconn

import (
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/usbdmx-golang/controller/enttec/dmxusbpro"
	log "github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

type DMXConn struct {
	dmx *dmxusbpro.EnttecDMXUSBProController
}

var dmxConn = DMXConn{}

func GetWriter() *DMXConn {
	return &dmxConn
}

// Initialize connection to the serial device over USB
func (d *DMXConn) Connect() {
	opts := options.GetAppOptions()
	port := opts.DmxWritePort
	baud := opts.DmxWriteBaudrate
	log.Infof("Opening DMX Serial using port %s", port)
	config := &serial.Config{Name: port, Baud: baud}

	// Create a controller and connect to it
	d.dmx = dmxusbpro.NewEnttecDMXUSBProController(config, true)
	if err := d.dmx.Connect(); err != nil {
		log.Fatalf("Failed to connect DMX Controller: %s", err)
	}
}

// Close connection to serial device (for cleanup)
func (d *DMXConn) Disconnect() {
	e := d.dmx.Disconnect()
	if e != nil {
		log.Fatal(e)
	}
}

// Set value for a given channel, remember to call 'Commit' to send them
func (d *DMXConn) Stage(channel int16, value byte) error {
	return d.dmx.Stage(channel, value)
}

// Send staged values to DMX
// If invalid staging occurred logs errors and resets
func (d *DMXConn) Commit() error {
	return d.dmx.Commit()
}

// Clear staged values
func (d *DMXConn) Clear() {
	d.dmx.ClearStage()
}
