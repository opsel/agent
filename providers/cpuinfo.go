package provider

import (
	"agent/config"
	"fmt"
	"log"
)

type CPUInfo struct{}

/**
* Gather information required by the module and pass
* them back to the main worker for submission
 */
func (provider CPUInfo) Gather() {

}

// WORKER
func (provider CPUInfo) Worker(cfg config.Config) {
	log.Printf("[INFO] MODULE: cpuinfo")

	/**
	* Check if the module is due for new submission
	* and then invoke the Gather function to gather
	* relevent informaiton
	 */
	fmt.Println(cfg.Agent.ID)

	/**
	 * Feed all the information to the upstream opsel
	 * server to process and store
	 */
	// http.Post(fmt.Sprintf("%s://%s/gather/", cfg.Server.Schema, cfg.Server.URI))
}
