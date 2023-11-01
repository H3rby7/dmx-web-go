package dmxbridge

import (
	dmxfader "github.com/H3rby7/dmx-web-go/internal/dmx/fader"
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
	writer       *dmxfader.FadingWriter
}

func NewDMXBridge(reader *dmxusbpro.EnttecDMXUSBProController, writer *dmxfader.FadingWriter) *DMXBridge {
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
	b.UpdateAll()
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
				b.writer.FadeTo(int16(k), v, dmxfader.FADE_IMMEDIATELY)
			}
		}
	}
}

// Update the writer with ALL stored values, if bridge is active
func (b *DMXBridge) UpdateAll() {
	if !b.isActive {
		log.Debugf("Skipping update, because bridge is not active")
		return
	}
	log.Debugf("Updating with %v", b.foreignInput)
	for i := range b.foreignInput {
		b.writer.FadeTo(int16(i), b.foreignInput[i], dmxfader.FADE_IMMEDIATELY)
	}
}
