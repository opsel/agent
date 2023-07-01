package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

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
		fmt.Println("[OPSEL] Running")

		fmt.Println(config.Agent.Modules)

		/**
		* Sleep the loop according to the time interval
		* defined in the opsel.yaml agent section
		 */
		time.Sleep(time.Duration(config.Agent.Interval) * time.Second)
	}

}
