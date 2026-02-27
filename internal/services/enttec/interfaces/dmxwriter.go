// Package enttecinterfaces defines service interfaces for these enttec services
package enttecinterfaces

// DMX Writer that takes care of fading channels to the desired values over time.
type EnttecDMXWriter interface {
	Write(values map[int16]byte)
	ConnectDMX()
	DisconnectDMX()
}
