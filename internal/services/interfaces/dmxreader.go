// Package interfaces defines service interfaces to interact with DMX
package interfaces

// DMXReader handles reading from DMX
type DMXReader interface {
	OnDMXChange(c chan map[int]byte)
	ConnectDMX()
	DisconnectDMX()
}
