package provider

import (
	"agent/config"
	"agent/utils"
	"database/sql"
	"log"
	"os"
)

type (
	System struct {
		Hostname string `json:"hostname"`
	}
)

/**
* Gather information required by the module and pass
* them back to the main worker for submission
 */
func (provider System) Gather() (System, error) {

	/**
	* Gather hostname of the system and include them in
	* the system struct
	 */
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("[ERROR] %s", err)
		return System{}, err
	}

	/**
	* [TODO]
	* Gather information about the network interfaces and
	* include them in the system struct
	 */

	// RETURN
	return System{
		Hostname: hostname,
	}, nil
}

// WORKER
func (provider System) Worker(cfg config.Config, db *sql.DB) {

	/**
	* Check if the module is due for new submission
	* and then invoke the Gather function to gather
	* relevent informaiton
	 */
	if isDue := utils.IsDue(db, "system", 10800); isDue {
		log.Printf("[INFO] MODULE: system")

		system, err := provider.Gather()
		if err != nil {
			log.Printf("[ERROR] %s", err)
			return
		}

		// Feeder
		if err := utils.Feeder(cfg, "/system", system); err != nil {
			log.Printf("[ERROR] %s", err)
			return
		}
	}
}
