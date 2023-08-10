package provider

import (
	"agent/config"
	"database/sql"
	"log"
)

type (
	Memory struct{}
)

/**
* Gather information required by the module and pass
* them back to the main worker for submission
 */
func (provider Memory) Gather() {}

// WORKER
func (provider Memory) Worker(cfg config.Config, db *sql.DB) {
	log.Printf("[INFO] MODULE: memory")

	/**
	* Check if the module is due for new submission
	* and then invoke the Gather function to gather
	* relevent informaiton
	 */

	/**
	 * Feed all the information to the upstream opsel
	 * server to process and store
	 */
	// http.Post(fmt.Sprintf("%s://%s/gather/", cfg.Server.Schema, cfg.Server.URI))
}
