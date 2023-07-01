package main

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
