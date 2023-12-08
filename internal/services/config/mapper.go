package config

import (
	models_chase "github.com/H3rby7/dmx-web-go/internal/model/chase"
	models_config "github.com/H3rby7/dmx-web-go/internal/model/config"
	models_dmx "github.com/H3rby7/dmx-web-go/internal/model/dmx"
	models_scene "github.com/H3rby7/dmx-web-go/internal/model/scene"
	models_trigger "github.com/H3rby7/dmx-web-go/internal/model/trigger"
)

// Map DMXValueForChannel
//
// From CONFIG models to DMX models
func mapDmxValueForChannel(in models_config.DMXValueForChannel) models_dmx.DMXValueForChannel {
	return models_dmx.DMXValueForChannel{
		Channel: in.Channel,
		Value:   in.Value,
	}
}

// Map Scene
//
// From CONFIG models to SCENE models
func mapScene(in models_config.Scene) models_scene.Scene {
	list := make([]models_dmx.DMXValueForChannel, 0, len(in.List))
	for _, v := range in.List {
		list = append(list, mapDmxValueForChannel(v))
	}
	return models_scene.Scene{
		List: list,
	}
}

// Map Step
//
// From CONFIG models to CHASE models
func mapStep(in models_config.Step) models_chase.Step {
	return models_chase.Step{
		DelayTimeMillis: in.DelayTimeMillis,
		BridgeActive:    in.BridgeActive,
		FadeTimeMillis:  in.FadeTimeMillis,
		Scene:           mapScene(in.Scene),
	}
}

// Map Chase
//
// From CONFIG models to CHASE models
func mapChase(in models_config.Chase) models_chase.Chase {
	steps := make([]models_chase.Step, 0, len(in.Chase))
	for _, v := range in.Chase {
		steps = append(steps, mapStep(v))
	}
	return models_chase.NewChase(in.Name, steps)
}

// Map Chases Array
//
// From CONFIG models to CHASE models
func mapChases(in []models_config.Chase) []models_chase.Chase {
	chases := make([]models_chase.Chase, 0, len(in))
	for _, v := range in {
		chases = append(chases, mapChase(v))
	}
	return chases
}

// Map Trigger
//
// From CONFIG models to TRIGGER models
func mapTrigger(in models_config.Trigger) models_trigger.Trigger {
	return models_trigger.Trigger{
		Source: in.Source,
		Goal:   in.Goal,
		Target: in.Target,
	}
}

// Map Trigger Array
//
// From CONFIG models to TRIGGER models
func mapTriggers(in []models_config.Trigger) []models_trigger.Trigger {
	triggers := make([]models_trigger.Trigger, 0, len(in))
	for _, v := range in {
		triggers = append(triggers, mapTrigger(v))
	}
	return triggers
}
