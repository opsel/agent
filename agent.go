package main

import (
	"agent/config"
	provider "agent/providers"
	"fmt"
	"log"
	"time"
)

/**
* Define the interface for the Provider to later use
* in seperate providers to implement
 */
type Provider interface {
	Worker(cfg config.Config)
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

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("[ERROR] %s", err.Error())
	}

	/** [TODO]
	* Check if the Agent ID is configured or not before
	* starting the modules otherwise upstream will get
	* lot of unwanted data feeds
	 */
	fmt.Println(cfg.Agent.ID)

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
		for _, module := range cfg.Agent.Modules {
			var provider Provider = ProviderFactory[module]
			go provider.Worker(cfg)
		}

		/**
		* Sleep the loop according to the time interval
		* defined in the opsel.yaml agent section
		 */
		time.Sleep(time.Duration(cfg.Agent.Interval) * time.Second)
	}

}
