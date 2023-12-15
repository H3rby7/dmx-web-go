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

const DEV_CONFIG_FILE_PATH = "./assets/sample-config.yaml"

// Attempt to load from multiple locations
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

// Attempt loading from location as defined by the ENV var
func loadConfig() (file models_config.ConfigFile, success bool) {
	opts := options.GetAppOptions()
	path := opts.ConfigFile
	contextLogger := log.WithField("filename", path)
	contextLogger.Infof("Attempting to read actions as defined by 'config' flag")
	return loadTriggersFromRelativePath(path)
}

// Loads actions from '../../assets/sample-actions.yaml'
// Fallback, or: When running in a development environment
func loadDevConfig() (file models_config.ConfigFile, success bool) {
	log.WithField("filename", DEV_CONFIG_FILE_PATH).Warn("Falling back to DEV Triggers ")
	return loadTriggersFromRelativePath(DEV_CONFIG_FILE_PATH)
}

func loadTriggersFromRelativePath(relPath string) (file models_config.ConfigFile, success bool) {
	absPath, err := filepath.Abs(relPath)
	if err != nil {
		log.WithField("filename", relPath).Error("Could resolve file ", err)
		return file, false
	}
	log.Debugf("Expanded relative path '%s' to absolute path '%s'", relPath, absPath)
	return loadTriggersFromAbsolutePath(absPath)
}

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
