// Package config allows using a yaml file to define [Trigger]s, [Chase]s and [Event]s.
package config

import (
	"io"
	"os"
	"path/filepath"

	models_config "github.com/H3rby7/dmx-web-go/internal/model/config"
	"github.com/H3rby7/dmx-web-go/internal/options"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// Attempt to load the config file as defined by the 'config' FLAG
func loadConfigFile() (config models_config.ConfigFile) {
	log.Infof("Loading configuration... ")
	config, success := loadConfig()
	if success {
		ValidateFile(config)
		return
	}
	var errorMessage string = `Could not load config! The application will look in the following places and take the first valid file:
		1. As specified by the flag 'config'`
	log.Error(errorMessage)
	panic(errorMessage)
}

// loadConfig attempts to load the config from location as defined by the 'config' FLAG
func loadConfig() (file models_config.ConfigFile, success bool) {
	opts := options.GetAppOptions()
	path := opts.ConfigFile
	contextLogger := log.WithField("filename", path)
	contextLogger.Infof("Attempting to read actions as defined by 'config' flag")
	return loadFromRelativePath(path)
}

// loadFromRelativePath takes a relative filepath and attempts to load the file as config file.
//
// Utilizes [loadFromAbsolutePath]
func loadFromRelativePath(relPath string) (file models_config.ConfigFile, success bool) {
	absPath, err := filepath.Abs(relPath)
	if err != nil {
		log.WithField("filename", relPath).Error("Could resolve file ", err)
		return file, false
	}
	log.Debugf("Expanded relative path '%s' to absolute path '%s'", relPath, absPath)
	return loadFromAbsolutePath(absPath)
}

// loadFromAbsolutePath takes an absolute filepath and attempts to load the file as config file.
func loadFromAbsolutePath(absPath string) (file models_config.ConfigFile, success bool) {
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
