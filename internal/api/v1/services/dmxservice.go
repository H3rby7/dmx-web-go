package apiv1services

import (
	models "github.com/H3rby7/dmx-web-go/internal/api/v1/models"
	"github.com/H3rby7/dmx-web-go/internal/dmx"
	"github.com/H3rby7/dmx-web-go/internal/options"
	log "github.com/sirupsen/logrus"
)

/*
Apply DMX values immediately to multiple channels
*/
func SetScene(data models.Scene, fadeDurationMillis int64) {
	opts := options.GetAppOptions()
	if ok, objection := opts.CanWriteDMX(); !ok {
		log.Warnf("%s -> Cannot set scene", objection)
		return
	}
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
	opts := options.GetAppOptions()
	if ok, objection := opts.CanWriteDMX(); !ok {
		log.Warnf("%s -> Cannot clear all", objection)
		return
	}
	dmx.GetFader().ClearAll()
}
