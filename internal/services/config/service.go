package config

import (
	models_chase "github.com/H3rby7/dmx-web-go/internal/model/chase"
	models_config "github.com/H3rby7/dmx-web-go/internal/model/config"
	models_event "github.com/H3rby7/dmx-web-go/internal/model/event"
	models_trigger "github.com/H3rby7/dmx-web-go/internal/model/trigger"
	log "github.com/sirupsen/logrus"
)

// Service holding the config data read from file
type ConfigService struct {
	config models_config.ConfigFile
}

/*
Create a new config service

Will read the config from the filesystem
*/
func NewConfigService() *ConfigService {
	log.Debugf("Creating new ConfigService")
	conf := loadConfigFile()
	return &ConfigService{
		config: conf,
	}
}

// Get the chases mapped to the internal model
func (svc *ConfigService) GetChases() []models_chase.Chase {
	return mapChases(svc.config.Chases)
}

// Get the triggers mapped to the internal model
func (svc *ConfigService) GetTriggers() []models_trigger.Trigger {
	return mapTriggers(svc.config.Triggers)
}

// Get the event sequences mapped to the internal model
func (svc *ConfigService) GetEventSequences() []models_event.EventSequence {
	return mapEventSequences(svc.config.EventSequences)
}
