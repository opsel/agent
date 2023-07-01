package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		Agent struct {
			Interval uint64   `yaml:"interval"`
			Modules  []string `yaml:"modules"`
		} `yaml:"agent"`

		Server struct {
			Schema string `yaml:"schema"`
			URI    string `yaml:"uri"`
		} `yaml:"server"`
	}
)

func NewConfig() (Config, error) {

	/**
	* Check for configuration file for agent to load
	* configuration. File lookup order will be as
	* following.
	*    - ./opsel.yaml
	*    - /etc/opsel/opsel.yaml
	 */
	var filename string
	if _, err := os.Stat("./opsel.yaml"); err != nil {
		if _, err := os.Stat("/etc/opsel/opsel.yaml"); err != nil {
			return Config{}, err
		} else {
			filename = "/etc/opsel/opsel.yaml"
		}
	} else {
		filename = "./opsel.yaml"
	}

	/**
	* Open configuration file found in the recent stage
	* and then parse with yaml.v3 parser
	 */
	file, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	/**
	* Read configuration file and parse the yaml file
	* according to the config.go struct
	 */
	payload, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := yaml.Unmarshal(payload, &config); err != nil {
		return Config{}, err
	}

	return config, nil

}
