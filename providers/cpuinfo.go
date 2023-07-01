package provider

import (
	"agent/config"
	"log"
)

type CPUInfo struct{}

func (provider CPUInfo) isDue() {

	/**
	*
	 */

}

func (provider CPUInfo) Worker(cfg config.Config) {
	log.Printf("[INFO] MODULE: cpuinfo")
}
