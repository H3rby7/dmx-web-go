// Package config allows using a yaml file to define [Trigger]s, [Chase]s and [Event]s.
package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	models_config "github.com/H3rby7/dmx-web-go/internal/model/config"
	"github.com/H3rby7/dmx-web-go/internal/options"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// DEV_CONFIG_FILE_PATH is the relative path of an example config for DEV purposes.
// This is only a fallback
const DEV_CONFIG_FILE_PATH = "./assets/sample-config.yaml"

// Attempt to load the config file (from multiple locations)
//
// Starts with the path defined by the FLAG.
// Fallback is defined by [DEV_CONFIG_FILE_PATH]
func loadConfigFile() (config models_config.ConfigFile) {
	log.Infof("Loading configuration... ")
	config, success := loadConfig()
	if success {
		ValidateFile(config)
		return
	}
	// Room for other loaders in order of precendence
	config, success = loadDevConfig()
	if success {
		ValidateFile(config)
		return
	}
	var errorMessage string = fmt.Sprintf(
		`Could not load config! The application will look in the following places and take the first valid file:
		1. As specified by the flag 'config',
		2. '%s' (the assets directory for development)`,
		DEV_CONFIG_FILE_PATH)
	log.Error(errorMessage)
	panic(errorMessage)
}

// loadConfig attempts to load the config from location as defined by the FLAG
func loadConfig() (file models_config.ConfigFile, success bool) {
	opts := options.GetAppOptions()
	path := opts.ConfigFile
	contextLogger := log.WithField("filename", path)
	contextLogger.Infof("Attempting to read actions as defined by 'config' flag")
	return loadTriggersFromRelativePath(path)
}

// loadDevConfig attempts to load the config from [DEV_CONFIG_FILE_PATH]
func loadDevConfig() (file models_config.ConfigFile, success bool) {
	log.WithField("filename", DEV_CONFIG_FILE_PATH).Warn("Falling back to DEV Triggers ")
	return loadTriggersFromRelativePath(DEV_CONFIG_FILE_PATH)
}

// loadTriggersFromRelativePath takes a relative filepath and attempts to load the file as config file.
//
// Utilizes [loadTriggersFromAbsolutePath]
func loadTriggersFromRelativePath(relPath string) (file models_config.ConfigFile, success bool) {
	absPath, err := filepath.Abs(relPath)
	if err != nil {
		log.WithField("filename", relPath).Error("Could resolve file ", err)
		return file, false
	}
	log.Debugf("Expanded relative path '%s' to absolute path '%s'", relPath, absPath)
	return loadTriggersFromAbsolutePath(absPath)
}

// loadTriggersFromRelativePath takes an absolute filepath and attempts to load the file as config file.
func loadTriggersFromAbsolutePath(absPath string) (file models_config.ConfigFile, success bool) {
	contextLogger := log.WithField("filename", absPath)
	contextLogger.Info("Loading actions ")

	fileHandle, err := os.Open(absPath)
	if err != nil {
		contextLogger.Error("Could not open file ", err)
		return file, false
	}
	// defer the closing of our file so that we can parse it later on
	defer fileHandle.Close()

	contextLogger.Debug("Successfully opened file ")

	// read our opened file as a byte array.
	byteValue, err := io.ReadAll(fileHandle)
	if err != nil {
		contextLogger.Error("Failed Reading File ", err)
		return file, false
	}

	// Unmarshall into  struct
	err = yaml.Unmarshal(byteValue, &file)
	if err != nil {
		contextLogger.Error("Failed YAML unmarshalling ", err)
		return file, false
	}

	contextLogger.Debug("Successfully unmarshalled YAML into struct")
	return file, true
}
