package utils

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

/**
* If the function return true if the module is due for
* execution and otherwise it'll return false
 */
func IsDue(db *sql.DB, module string, delta int64) bool {

	/**
	* Check if the config for the last submission is configured
	* and if not create the value with current unix timestamp
	 */
	var module_exec_on int64
	if err := db.QueryRow("SELECT `value` FROM `config` WHERE `name` = ?", fmt.Sprintf("module_%s_exec_on", module)).Scan(&module_exec_on); err != nil {
		log.Printf("[ERROR] VALIDATOR::IsDue %s", err.Error())

		/**
		* Insert the record with current time because we don't
		* have any record regarding this module right now.
		 */
		if _, err := db.Exec("INSERT INTO `config` ('name', 'value') VALUES (?, ?)", fmt.Sprintf("module_%s_exec_on", module), time.Now().Unix()); err != nil {
			log.Printf("[ERROR] VALIDATOR::IsDue %s", err.Error())
			return false
		} else {
			return true
		}

	}

	if (time.Now().Unix() - module_exec_on) > delta {

		/**
		* Update database with the new time delta and return
		* true value to execute the module
		 */
		if _, err := db.Exec("UPDATE `config` SET `value` = ? WHERE `name` = ?", time.Now().Unix(), fmt.Sprintf("module_%s_exec_on", module)); err != nil {
			log.Printf("[ERROR] VALIDATOR::IsDue %s", err.Error())
			return false
		} else {
			return true
		}

	} else {
		return false
	}

}
