package options

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type DMXWebOptions struct {
	HttpPort   string
	LogLevel   log.Level
	SerialPort string
}

// Local instance holding our settings
var optionsInstance = DMXWebOptions{
	HttpPort:   getServerPort(),
	LogLevel:   getLogLevel(),
	SerialPort: getSerialPort(),
}

/*
Get Options
*/
func GetDMXWebOptions() DMXWebOptions {
	return optionsInstance
}

/*
Take port from env 'HTTP_SERVER_PORT'
Default: 8080
*/
func getServerPort() string {
	port := os.Getenv("HTTP_SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

/*
Take log level from env 'LOG_LEVEL'
Default: Info
*/
func getLogLevel() log.Level {
	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}
	return logLevel
}

/*
Take serial port from env 'SERIAL_PORT'
Default: 4
*/
func getSerialPort() string {
	port := os.Getenv("SERIAL_PORT")
	if port == "" {
		port = "COM6"
	}
	return port
}
