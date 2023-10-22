package dmxbridge

import (
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/usbdmx-golang/controller/enttec/dmxusbpro"
	"github.com/H3rby7/usbdmx-golang/controller/enttec/dmxusbpro/messages"
	log "github.com/sirupsen/logrus"
)

type DMXBridge struct {
	// is bridging active?
	isActive bool
	// Holds DMX data, as DMX starts with channel '1' the index '0' is unused.
	foreignInput []byte
	reader       *dmxusbpro.EnttecDMXUSBProController
	writer       *dmxusbpro.EnttecDMXUSBProController
}

func NewDMXBridge(reader *dmxusbpro.EnttecDMXUSBProController, writer *dmxusbpro.EnttecDMXUSBProController) *DMXBridge {
	opts := options.GetAppOptions()
	if ok, objection := opts.CanBridge(); !ok {
		log.Panicf("%s, cannot bridge.", objection)
	}
	channels := opts.DmxChannelCount
	b := &DMXBridge{
		isActive:     false,
		foreignInput: make([]byte, channels+1),
		reader:       reader,
		writer:       writer,
	}
	return b
}

// Start passing on any data that is read
func (b *DMXBridge) Activate() {
	b.isActive = true
	b.Update()
}

// Stop passing on data that is read
func (b *DMXBridge) Deactivate() {
	b.isActive = false
}

// Configure reader to 'changes only' and activate the bridge
func (b *DMXBridge) BridgeDMX() {
	b.reader.SwitchReadMode(1)
	c := make(chan messages.EnttecDMXUSBProApplicationMessage)
	go b.reader.OnDMXChange(c, 15)
	for msg := range c {
		cs, err := messages.ToChangeSet(msg)
		if err != nil {
			log.Printf("Could not convert to changeset, but read \tlabel=%v \tdata=%v", msg.GetLabel(), msg.GetPayload())
		} else {
			for k, v := range cs {
				b.foreignInput[k] = v
			}
			b.Update()
		}
	}
}

// Update the writer with stored values, if bridge is active
func (b *DMXBridge) Update() {
	if !b.isActive {
		log.Debugf("Skipping update, because bridge is not active")
		return
	}
	log.Debugf("Updating with %v", b.foreignInput)
	for i := range b.foreignInput {
		b.writer.Stage(int16(i), b.foreignInput[i])
	}
	b.writer.Commit()
}
