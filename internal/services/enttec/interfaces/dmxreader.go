// Package enttecinterfaces defines service interfaces for these enttec services
package enttecinterfaces

// DMXReader handles reading from DMX
type EnttecDMXReader interface {
	OnDMXChange(c chan map[int]byte)
	ConnectDMX()
	DisconnectDMX()
}
