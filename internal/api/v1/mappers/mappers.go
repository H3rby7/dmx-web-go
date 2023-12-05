package apiv1mappers

import (
	apiv1dtos "github.com/H3rby7/dmx-web-go/internal/api/v1/dtos"
	models_dmx "github.com/H3rby7/dmx-web-go/internal/model/dmx"
	models_scene "github.com/H3rby7/dmx-web-go/internal/model/scene"
)

/*
Map DMXValueForChannel

From API V1 DTO models to DMX models
*/
func MapDmxValueForChannel(in apiv1dtos.DMXValueForChannel) models_dmx.DMXValueForChannel {
	return models_dmx.DMXValueForChannel{
		Channel: in.Channel,
		Value:   in.Value,
	}
}

/*
Map Scene

From API V1 DTO models to SCENE models
*/
func MapScene(in apiv1dtos.Scene) models_scene.Scene {
	list := make([]models_dmx.DMXValueForChannel, 0, len(in.List))
	for _, v := range in.List {
		list = append(list, MapDmxValueForChannel(v))
	}
	return models_scene.Scene{
		List: list,
	}
}
