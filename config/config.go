package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Host          string   `json:"host"`
	Port          string   `json:"port"`
	Https         bool     `json:"https"`
	Logs          bool     `json:"logs"`
	AuthorizedIPs []string `json:"authorized_ips"`
}

func GetConfig() (*Config, error) {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		config := &Config{}

		config.Host = "localhost"
		config.Port = "80"
		config.Https = false
		config.Logs = true
		config.AuthorizedIPs = []string{}

		configBytes, _ := json.Marshal(config)
		var configBuffer bytes.Buffer
		_ = json.Indent(&configBuffer, configBytes, "", "\t")

		if err = ioutil.WriteFile("config.json", configBuffer.Bytes(), 0777); err != nil {
			return nil, err
		}
	}

	data, err := ioutil.ReadFile("config.json")

	if err != nil {
		return nil, err
	}

	config := &Config{}

	json.Unmarshal(data, config)

	return config, nil
}
