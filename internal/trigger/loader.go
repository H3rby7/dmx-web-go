package trigger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const ENV_NAME_PATH = "DMX_WEB_ACTIONS_PATH"
const ACTIONS_FILE_NAME = "actions.yaml"
const DEV_ASSETS_ACTIONS_FILE_PATH = "./assets/sample-actions.yaml"

// Attempt to load from multiple locations
func Load() (actions ActionsFile) {
	log.Infof("Loading actions... ")
	actions, success := loadTriggersFromEnvPath()
	if success {
		ValidateFile(actions)
		return
	}
	actions, success = loadFromExecDir()
	if success {
		ValidateFile(actions)
		return
	}
	// Room for other loaders in order of precendence
	actions, success = loadDevTriggers()
	if success {
		ValidateFile(actions)
		return
	}
	var errorMessage string = fmt.Sprintf(
		`Could not load actions! The application will look in the following places and take the first valid file:
		1. '%s' as specified by the environment variable,
		2. '%s' file next to the binary, 
		3. '%s' (the assets directory for development)`,
		ENV_NAME_PATH, ACTIONS_FILE_NAME, DEV_ASSETS_ACTIONS_FILE_PATH)
	log.Error(errorMessage)
	panic(errorMessage)
}

// Attempt loading from location as defined by the ENV var
func loadTriggersFromEnvPath() (file ActionsFile, success bool) {
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
func loadFromExecDir() (file ActionsFile, success bool) {
	log.WithField("filename", ACTIONS_FILE_NAME).Info("Reading actions from exec directory ")
	_, err := os.Stat(ACTIONS_FILE_NAME)
	if err != nil {
		log.Debug(fmt.Sprintf("No '%s' file present in dir of executable ", ACTIONS_FILE_NAME))
		return file, false
	}
	return loadTriggersFromRelativePath(ACTIONS_FILE_NAME)
}

// Loads actions from '../../assets/sample-actions.yaml'
// Fallback, or: When running in a development environment
func loadDevTriggers() (file ActionsFile, success bool) {
	log.WithField("filename", DEV_ASSETS_ACTIONS_FILE_PATH).Warn("Falling back to DEV Triggers ")
	return loadTriggersFromRelativePath(DEV_ASSETS_ACTIONS_FILE_PATH)
}

func loadTriggersFromRelativePath(relPath string) (file ActionsFile, success bool) {
	absPath, err := filepath.Abs(relPath)
	if err != nil {
		log.WithField("filename", relPath).Error("Could resolve file ", err)
		return file, false
	}
	log.Debugf("Expanded relative path '%s' to absolute path '%s'", relPath, absPath)
	return loadTriggersFromAbsolutePath(absPath)
}

func loadTriggersFromAbsolutePath(absPath string) (file ActionsFile, success bool) {
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
