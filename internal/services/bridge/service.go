package bridge

import (
	"github.com/H3rby7/dmx-web-go/internal/dmx"
)

// BridgeService handles the DMX Bridge
type BridgeService struct{}

func NewBridgeService() *BridgeService {
	return &BridgeService{}
}

// Activate activates the DMX bridge
func (b *BridgeService) Activate() {
	dmx.GetBridge().Activate()
}

// Activate deactivates the DMX bridge
func (b *BridgeService) Deactivate() {
	dmx.GetBridge().Deactivate()
}
