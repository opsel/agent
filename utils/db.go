package utils

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func DB() *sql.DB {

	/**
	* Check and create the database file using OS library
	* and once we have the file we can continue database
	* initiation process
	 */
	if _, err := os.Stat("/usr/share/opsel/agent.db"); err != nil {
		if _, err := os.Create("/usr/share/opsel/agent.db"); err != nil {
			log.Fatalf("[ERROR] DB: %s", err.Error())
		}
	}

	/**
	* Database file is already exists and we can try establishing
	* connection to the database.
	 */
	db, err := sql.Open("sqlite3", "/usr/share/opsel/agent.db")
	if err != nil {
		log.Fatalf("[ERROR] DB: %s", err.Error())
	}

	/**
	* Connection is established and now we execute the clean-up
	* activities to the database and create data tables as we
	* required.
	 */
	if stmt, err := db.Prepare(`DROP TABLE IF EXISTS config`); err != nil {
		log.Fatalf("[ERROR] DB: %s", err.Error())
	} else {
		if _, err := stmt.Exec(); err != nil {
			log.Fatalf("[ERROR] DB: %s", err.Error())
		} else {
			log.Printf("[ERROR] DB: DROP TABLE `config`")
		}
	}

	/**
	* Populate config database table with required columns
	* so we can use it in the future to operations
	 */
	if stmt, err := db.Prepare(`CREATE TABLE config ('id' INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 'name' TEXT NOT NULL UNIQUE, 'value' TEXT NOT NULL);`); err != nil {
		log.Fatalf("[ERROR] DB: %s", err.Error())
	} else {
		if _, err := stmt.Exec(); err != nil {
			log.Fatalf("[ERROR] DB: %s", err.Error())
		} else {
			log.Printf("[ERROR] DB: POPULATE TABLE `config`")
		}
	}

	/**
	* Finally everything seems fine and we are ready to return
	* database instance back
	 */
	return db

}
