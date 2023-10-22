package apiv1

import (
	dmxconn "github.com/H3rby7/dmx-web-go/internal/dmx"
)

// apply dmx values to multiple channels immediately
func dmxSetMulti(data MultipleDMXValueForChannel) {
	dmx := dmxconn.GetFader()
	for _, entry := range data.List {
		dmx.FadeTo(entry.Channel, entry.Value, 0)
	}
}
