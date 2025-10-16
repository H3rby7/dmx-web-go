// Package bridge defines the DMX [BridgeService] as a tool to forward READ DMX values to the WRITER
package bridge

import (
	models_fader "github.com/H3rby7/dmx-web-go/internal/model/fader"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/dmx-web-go/internal/services/fading"
	"github.com/H3rby7/dmx-web-go/internal/services/reader"
	log "github.com/sirupsen/logrus"
)

// BridgeService takes care of bridging DMX values from READ to WRITE
type BridgeService struct {
	// is bridging active?
	isActive bool
	// Holds DMX data, as DMX starts with channel '1' the index '0' is unused.
	foreignInput  []byte
	reader        reader.DMXReader
	fadingService fading.DMXFader
}

// NewBridgeService creates a new [BridgeService] instance with proper defaults
func NewBridgeService(reader reader.DMXReader, fadingService fading.DMXFader) *BridgeService {
	log.Debugf("Creating new BridgeService")
	opts := options.GetAppOptions()
	channels := opts.DmxChannelCount
	b := &BridgeService{
		isActive:      false,
		foreignInput:  make([]byte, channels+1),
		reader:        reader,
		fadingService: fadingService,
	}

	if ok, objection := opts.CanBridge(); ok {
		b.Activate(models_fader.FADE_IMMEDIATELY)
		go b.bridgeDMX()
	} else {
		log.Infof("%s -> Skipping to bridge.", objection)
	}

	return b
}

// Activate the DMX bridge
//
// This enables passing on any data that is read
func (b *BridgeService) Activate(fadeDurationMillis int64) {
	if b.isActive {
		log.Tracef("Bridge already active")
		return
	}
	b.isActive = true
	log.Infof("Activating bridge over %v millis", fadeDurationMillis)
	opts := options.GetAppOptions()
	if ok, objection := opts.CanBridge(); !ok {
		log.Infof("%s -> Skipping updateAll", objection)
	} else {
		b.updateAll(fadeDurationMillis)
	}
}

// Deactivate the DMX bridge
//
// This stops passing on data that is read
func (b *BridgeService) Deactivate(fadeDurationMillis int64) {
	if !b.isActive {
		log.Tracef("Bridge already inactive")
		return
	}
	log.Infof("Deactivating bridge of %v millis", fadeDurationMillis)
	b.isActive = false
	b.clearOutput(fadeDurationMillis)
}

// Register On-DMX-Change Channel
func (b *BridgeService) bridgeDMX() {
	c := make(chan map[int]byte)
	go b.reader.OnDMXChange(c)
	for cs := range c {
		for k, v := range cs {
			b.foreignInput[k] = v
			b.fadingService.FadeTo(int16(k), v, models_fader.FADE_IMMEDIATELY)
		}
	}
}

// Update the writer with ALL stored values, if bridge is active
func (b *BridgeService) updateAll(fadeDurationMillis int64) {
	if !b.isActive {
		log.Debugf("Skipping update, because bridge is not active")
		return
	}
	log.Debugf("Updating over %v millis with %v", fadeDurationMillis, b.foreignInput)
	for i := range b.foreignInput {
		b.fadingService.FadeTo(int16(i), b.foreignInput[i], fadeDurationMillis)
	}
}

// Clear Bridge Output over time
//
// Fade any channel, that is != 0 to 0
func (b *BridgeService) clearOutput(fadeDurationMillis int64) {
	log.Infof("Clearing bridge output over %v millis", fadeDurationMillis)
	for k, v := range b.foreignInput {
		if v != 0 {
			b.fadingService.FadeTo(int16(k), v, fadeDurationMillis)
		}
	}
}
