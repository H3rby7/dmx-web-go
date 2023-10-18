package setup

import dmxconn "github.com/H3rby7/dmx-web-go/internal/dmx"

// Configure DMX connections with respect to the app options
func SetUpDMX() {
	dmxconn.GetWriter().Connect()
}
