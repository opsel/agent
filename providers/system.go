package provider

import (
	"agent/config"
	"agent/utils"
	"database/sql"
	"log"
)

type (
	System struct{}
)

/**
* Gather information required by the module and pass
* them back to the main worker for submission
 */
func (provider System) Gather() {}

// WORKER
func (provider System) Worker(cfg config.Config, db *sql.DB) {

	/**
	* Check if the module is due for new submission
	* and then invoke the Gather function to gather
	* relevent informaiton
	 */
	if isDue := utils.IsDue("system", db); isDue == true {
		log.Printf("[INFO] MODULE: system is due for execution")
	} else {
		log.Println("[INFO] MODULE: system is not due for execution")
	}

	/**
	 * Feed all the information to the upstream opsel
	 * server to process and store
	 */
	// http.Post(fmt.Sprintf("%s://%s/gather/", cfg.Server.Schema, cfg.Server.URI))
}
