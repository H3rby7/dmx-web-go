package dmxconn

import (
	"fmt"

	"github.com/akualab/dmx"
	log "github.com/sirupsen/logrus"
)

type DMXConn struct {
	dmx  *dmx.DMX
	errs []error
}

var dmxConn = DMXConn{}

func GetDMXConn() *DMXConn {
	return &dmxConn
}

// Initialize connection to the serial device over USB
func (d *DMXConn) Connect(port string) {
	log.Infof("Opening DMX Serial using port %s", port)
	dmx, e := dmx.NewDMXConnection(port)
	if e != nil {
		log.Fatal(e)
	}
	d.dmx = dmx
}

// Close connection to serial device (for cleanup)
func (d *DMXConn) Disconnect() {
	e := d.dmx.Close()
	if e != nil {
		log.Fatal(e)
	}
}

// Set value for a given channel, remember to call 'Commit' to send them
func (d *DMXConn) Stage(channel int, value int) (ok bool) {
	ok = true
	if err := checkChannel(channel); err != nil {
		ok = false
		d.errs = append(d.errs, err)
	}
	if err := checkValue(value); err != nil {
		ok = false
		d.errs = append(d.errs, err)
	}
	if ok {
		d.dmx.SetChannel(channel, byte(value))
	}
	return
}

// Send staged values to DMX
// If invalid staging occurred logs errors and resets
func (d *DMXConn) Commit() error {
	if len(d.errs) == 0 {
		return d.dmx.Render()
	}
	for _, e := range d.errs {
		log.Error(e)
	}
	d.Reset()
	return fmt.Errorf("staged commands had errors")
}

// Get current errors
func (d *DMXConn) GetErrors() []error {
	return d.errs
}

// Clear values, reset errors and 'ok'
func (d *DMXConn) Reset() {
	// To check: clear all really clearing all correctly?
	// Might be that with the next 'send' we lose all previously sent values in the DMX?
	d.dmx.ClearAll()
	d.errs = make([]error, 0)
}

// Send some static test values
func (d *DMXConn) Test() error {
	d.Stage(1, 100)
	d.Stage(2, 70)
	d.Stage(3, 130)
	d.Stage(4, 180)
	return d.Commit()
}
