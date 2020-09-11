package services

import (
	"encoding/json"
	"io/ioutil"
)

//Config holds our current application config.
type Config struct {
	DatabaseURI string `json:"database_uri"`
	APIPort     string `json:"api_port"`
}

//LoadConfigFromFile loads and unmarshales our config.
func LoadConfigFromFile(filepath string) (Config, error) {
	config := Config{}

	configBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
