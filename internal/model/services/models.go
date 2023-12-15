// Package models_services provides structs for simplified dependency injection
package models_services

import (
	"github.com/H3rby7/dmx-web-go/internal/services/bridge"
	"github.com/H3rby7/dmx-web-go/internal/services/chase"
	"github.com/H3rby7/dmx-web-go/internal/services/config"
	"github.com/H3rby7/dmx-web-go/internal/services/dmx"
	"github.com/H3rby7/dmx-web-go/internal/services/event"
	"github.com/H3rby7/dmx-web-go/internal/services/fading"
	"github.com/H3rby7/dmx-web-go/internal/services/trigger"
)

// ApplicationServices is a container to hold the services of the application
type ApplicationServices struct {
	BridgeService        *bridge.BridgeService
	ChaseService         *chase.ChaseService
	ConfigService        *config.ConfigService
	EventSequenceService *event.EventSequenceService
	FadingService        *fading.FadingService
	DMXReaderService     *dmx.DMXReaderService
	TriggerService       *trigger.TriggerService
}
