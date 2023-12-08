package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	models_config "github.com/H3rby7/dmx-web-go/internal/model/config"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const ENV_NAME_PATH = "DMX_WEB_CONFIG_PATH"
const CONFIG_FILE_NAME = "config.yaml"
const DEV_CONFIG_FILE_PATH = "./assets/sample-config.yaml"

// Attempt to load from multiple locations
func loadConfigFile() (config models_config.ConfigFile) {
	log.Infof("Loading configuration... ")
	config, success := loadTriggersFromEnvPath()
	if success {
		ValidateFile(config)
		return
	}
	config, success = loadFromExecDir()
	if success {
		ValidateFile(config)
		return
	}
	// Room for other loaders in order of precendence
	config, success = loadDevTriggers()
	if success {
		ValidateFile(config)
		return
	}
	var errorMessage string = fmt.Sprintf(
		`Could not load config! The application will look in the following places and take the first valid file:
		1. '%s' as specified by the environment variable,
		2. '%s' file next to the binary, 
		3. '%s' (the assets directory for development)`,
		ENV_NAME_PATH, CONFIG_FILE_NAME, DEV_CONFIG_FILE_PATH)
	log.Error(errorMessage)
	panic(errorMessage)
}

// Attempt loading from location as defined by the ENV var
func loadTriggersFromEnvPath() (file models_config.ConfigFile, success bool) {
	envPath, isset := os.LookupEnv(ENV_NAME_PATH)
	if !isset {
		log.Debugf("ENV '%s' is not set ", ENV_NAME_PATH)
		return file, false
	}
	contextLogger := log.WithField("filename", envPath)
	contextLogger.Infof("Attempting to read actions as defined by '%s' ", ENV_NAME_PATH)

	absPath, err := filepath.Abs(envPath)
	if err != nil {
		contextLogger.Error("Could resolve file ", err)
		return file, false
	}
	return loadTriggersFromAbsolutePath(absPath)
}

// Loads from the yaml put next to executable
func loadFromExecDir() (file models_config.ConfigFile, success bool) {
	log.WithField("filename", CONFIG_FILE_NAME).Debug("Attempting to read from exec directory ")
	_, err := os.Stat(CONFIG_FILE_NAME)
	if err != nil {
		log.Debug(fmt.Sprintf("No '%s' file present in dir of executable ", CONFIG_FILE_NAME))
		return file, false
	}
	return loadTriggersFromRelativePath(CONFIG_FILE_NAME)
}

// Loads actions from '../../assets/sample-actions.yaml'
// Fallback, or: When running in a development environment
func loadDevTriggers() (file models_config.ConfigFile, success bool) {
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
