package bridge

import (
	models_fader "github.com/H3rby7/dmx-web-go/internal/model/fader"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/dmx-web-go/internal/services/dmx"
	"github.com/H3rby7/dmx-web-go/internal/services/fading"
	"github.com/H3rby7/usbdmx-golang/controller/enttec/dmxusbpro/messages"
	log "github.com/sirupsen/logrus"
)

// BridgeService handles the DMX Bridge
type BridgeService struct {
	// is bridging active?
	isActive bool
	// Holds DMX data, as DMX starts with channel '1' the index '0' is unused.
	foreignInput []byte
	reader       *dmx.DMXReaderService
	writer       *fading.FadingService
}

func NewBridgeService(reader *dmx.DMXReaderService, writer *fading.FadingService) *BridgeService {
	log.Debugf("Creating new BridgeService")
	opts := options.GetAppOptions()
	channels := opts.DmxChannelCount
	b := &BridgeService{
		isActive:     false,
		foreignInput: make([]byte, channels+1),
		reader:       reader,
		writer:       writer,
	}

	if ok, objection := opts.CanBridge(); ok {
		b.Activate()
		go b.bridgeDMX()
	} else {
		log.Infof("%s -> Skipping to bridge.", objection)
	}

	return b
}

// Activate the DMX bridge
//
// This enables passing on any data that is read
func (b *BridgeService) Activate() {
	if b.isActive {
		log.Tracef("Bridge already active")
		return
	}
	b.isActive = true
	b.updateAll()
}

// Deactivate the DMX bridge
//
// This stops passing on data that is read
func (b *BridgeService) Deactivate() {
	if !b.isActive {
		log.Tracef("Bridge already inactive")
		return
	}
	b.isActive = false
}

// Register On-DMX-Change Channel
func (b *BridgeService) bridgeDMX() {
	c := make(chan messages.EnttecDMXUSBProApplicationMessage)
	go b.reader.OnDMXChange(c)
	for msg := range c {
		cs, err := messages.ToChangeSet(msg)
		if err != nil {
			log.Warnf("Could not convert to changeset, but read \tlabel=%v \tdata=%v", msg.GetLabel(), msg.GetPayload())
		} else {
			for k, v := range cs {
				b.foreignInput[k] = v
				b.writer.FadeTo(int16(k), v, models_fader.FADE_IMMEDIATELY)
			}
		}
	}
}

// Update the writer with ALL stored values, if bridge is active
func (b *BridgeService) updateAll() {
	if !b.isActive {
		log.Debugf("Skipping update, because bridge is not active")
		return
	}
	log.Debugf("Updating with %v", b.foreignInput)
	for i := range b.foreignInput {
		b.writer.FadeTo(int16(i), b.foreignInput[i], models_fader.FADE_IMMEDIATELY)
	}
}
