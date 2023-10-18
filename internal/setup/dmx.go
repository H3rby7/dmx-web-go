package setup

import dmxconn "github.com/H3rby7/dmx-web-go/internal/dmx"

func SetUpDMX() {
	dmxconn.GetWriter().Connect()
}
