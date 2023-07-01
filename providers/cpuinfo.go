package provider

import (
	"fmt"
)

type CPUInfo struct{}

func (provider CPUInfo) isDue() {

	/**
	*
	 */

}

func (provider CPUInfo) Worker(schema string, uri string) {
	fmt.Println("CPUInfo::worker()")
}
