package main

import (
	provider "agent/providers"
	"io"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

/**
* Define the interface for the Provider to later use
* in seperate providers to implement
 */
type Provider interface {
	Worker(schema string, uri string)
}

/**
* Define list of available providers into the following
* factory var to utilize later in dynamic provider
* loading.
 */
var ProviderFactory map[string]Provider = map[string]Provider{
	"cpu_info": &provider.CPUInfo{},
	"mem_info": &provider.CPUInfo{},
}

func main() {

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
			log.Fatal("[ERROR] no configuration file found")
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
		log.Fatalf("[ERROR] %s", err.Error())
	}
	defer file.Close()

	/**
	* Read configuration file and parse the yaml file
	* according to the config.go struct
	 */
	payload, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("[ERROR] %s", err.Error())
	}

	var config Config
	if err := yaml.Unmarshal(payload, &config); err != nil {
		log.Fatalf("[ERROR] %s", err.Error())
	}

	/**
	* Loop each module to get the relevent information
	* and feed them into REST API endpoint configured
	* in the opsel.yaml configuration
	 */
	for {

		/**
		* Dispatch each provider in a different go routine
		* so each and every provider will do the job without
		* render blocking
		 */
		for _, module := range config.Agent.Modules {
			var provider Provider = ProviderFactory[module]
			go provider.Worker(config.Server.Schema, config.Server.URI)
		}

		/**
		* Sleep the loop according to the time interval
		* defined in the opsel.yaml agent section
		 */
		time.Sleep(time.Duration(config.Agent.Interval) * time.Second)
	}

}
