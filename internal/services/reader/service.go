// Package reader provides tools to READ from DMX
package reader

// DMXReader handles reading from DMX
type DMXReader interface {
	OnDMXChange(c chan map[int]byte)
	ConnectDMX()
	DisconnectDMX()
}
