package configuration

import (
	"encoding/json"
	"io/ioutil"
)

// Config contains the current configuration settings.
var Config Configuration

// ConfigFileLocation contains the location of the configuration file.
var ConfigFileLocation string

// Configuration stores the main configuration for the application.
type Configuration struct {
	Expiry int // The expiry of a login request in minutes. Defaults is 30.
}

// Read reads the configuration file and parses it. If any errors are encountered, a panic will occur.
func Read() {
	f, err := ioutil.ReadFile(ConfigFileLocation)

	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(f, &Config)
	if err != nil {
		panic(err.Error())
	}
}
