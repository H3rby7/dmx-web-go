package dmx

import (
	dmxbridge "github.com/H3rby7/dmx-web-go/internal/dmx/bridge"
	dmxfader "github.com/H3rby7/dmx-web-go/internal/dmx/fader"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/usbdmx-golang/controller/enttec/dmxusbpro"
)

var reader *dmxusbpro.EnttecDMXUSBProController
var writer *dmxusbpro.EnttecDMXUSBProController
var faderWriter *dmxfader.FadingWriter
var bridge *dmxbridge.DMXBridge

func GetFader() *dmxfader.FadingWriter {
	return faderWriter
}

func GetBridge() *dmxbridge.DMXBridge {
	return bridge
}

func Initialize() {
	opts := options.GetAppOptions()
	writer = createWriter()
	if ok, _ := opts.CanWriteDMX(); ok {
		faderWriter = dmxfader.NewFadingWriter(writer)
		faderWriter.Start()
	}
	reader = createReader()
	if opts.DmxBridge {
		bridge = dmxbridge.NewDMXBridge(reader, faderWriter)
		bridge.Activate()
		go bridge.BridgeDMX()
	}
}

func Shutdown() {
	shutdownReader(reader)
	shutdownWriter(writer)
	opts := options.GetAppOptions()
	if ok, _ := opts.CanWriteDMX(); ok {
		faderWriter.Stop()
	}
}
