package main

import (
	"agent/config"
	provider "agent/providers"
	"agent/utils"
	"database/sql"
	"log"
	"time"
)

/**
* Define the interface for the Provider to later use
* in seperate providers to implement
 */
type Provider interface {
	Worker(cfg config.Config, db *sql.DB)
}

/**
* Define list of available providers into the following
* factory var to utilize later in dynamic provider
* loading.
 */
var ProviderFactory map[string]Provider = map[string]Provider{
	"system":    &provider.System{},
	"processor": &provider.Processor{},
	"memory":    &provider.Memory{},
}

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("[ERROR] %s", err.Error())
	}

	/**
	* Implement a SQLite database to store runtime related
	* information about the agent.
	 */
	db := utils.DB()

	/**
	* Check if the Agent ID is configured or not before
	* starting the modules otherwise upstream will get
	* lot of unwanted data feeds
	 */
	if len(cfg.Agent.ID) != 36 {
		log.Fatal("[ERROR]: Agent ID is invalid")
	}

	/**
	* Loop each module to get the relevent information
	* and feed them into REST API endpoint configured
	* in the opsel.yaml configuration
	 */
	for {

		/**
		* Dispatch required modules without adding them to the
		* extra module set. These modules will always required
		* for the functionality.
		*     - System
		 */
		var SystemProvider Provider = ProviderFactory["system"]
		go SystemProvider.Worker(cfg, db)

		/**
		* Dispatch each provider in a different go routine
		* so each and every provider will do the job without
		* render blocking
		 */
		for _, module := range cfg.Agent.Modules {
			var provider Provider = ProviderFactory[module]
			go provider.Worker(cfg, db)
		}

		/**
		* Sleep the loop according to the time interval
		* defined in the opsel.yaml agent section
		 */
		time.Sleep(time.Duration(cfg.Agent.Interval) * time.Second)
	}

}
