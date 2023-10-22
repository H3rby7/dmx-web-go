package apiv1services

import (
	models "github.com/H3rby7/dmx-web-go/internal/api/v1/models"
	"github.com/H3rby7/dmx-web-go/internal/dmx"
	log "github.com/sirupsen/logrus"
)

/*
Apply DMX values immediately to multiple channels
*/
func SetScene(data models.Scene, fadeDurationMillis int64) {
	log.Debugf("Using a fade duration of %d millis, setting scene: %v", fadeDurationMillis, data.List)
	dmx := dmx.GetFader()
	for _, entry := range data.List {
		dmx.FadeTo(entry.Channel, entry.Value, fadeDurationMillis)
	}
}

/*
Set all DMX values to 0 immediately
*/
func ClearAll() {
	dmx.GetFader().ClearAll()
}
