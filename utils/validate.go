package utils

import (
	"database/sql"
	"fmt"
)

/**
* If the function return true if the module is due for
* execution and otherwise it'll return false
 */
func IsDue(module string, db *sql.DB) bool {

	res, err := db.Exec("INSERT INTO `config` ('name', 'value') VALUES ('module_system_exec_on', '1691644399')")
	fmt.Println(err)
	fmt.Println(res)

	/**
	* Check if the config for the last submission is configured
	* and if not create the value with current unix timestamp
	 */
	var module_system_exec_on string
	if err := db.QueryRow("SELECT 'value' FROM `config` WHERE 'name' = 'module_system_exec_on'").Scan(&module_system_exec_on); err != nil {
		fmt.Println(err)
	}
	fmt.Println(module_system_exec_on)
	return false
}
