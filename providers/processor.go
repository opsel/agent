package provider

import (
	"agent/config"
	"agent/utils"
	"database/sql"
	"log"

	"github.com/prometheus/procfs"
)

type (
	Core struct {
		ID        uint8    `json:"id"`
		ModelName string   `json:"name"`
		MHz       float64  `json:"mhz"`
		Flags     []string `json:"flags"`
	}
	Processor struct {
		Cores []Core `json:"cores"`
	}
)

/**
* Gather information required by the module and pass
* them back to the main worker for submission
 */
func (provider Processor) Gather() (Processor, error) {

	pfs, err := procfs.NewFS("/proc")
	if err != nil {
		return Processor{}, err
	}

	CPUs, err := pfs.CPUInfo()
	if err != nil {
		return Processor{}, err
	}

	var cores []Core
	for _, CPU := range CPUs {
		// fmt.Printf("INFO: %d\n", CPU.Processor)
		cores = append(cores, Core{
			ID:        uint8(CPU.Processor),
			ModelName: CPU.ModelName,
			MHz:       CPU.CPUMHz,
			Flags:     CPU.Flags,
		})
	}

	return Processor{
		Cores: cores,
	}, nil

}

// WORKER
func (provider Processor) Worker(cfg config.Config, db *sql.DB) {

	/**
	* Check if the module is due for new submission
	* and then invoke the Gather function to gather
	* relevent informaiton
	 */
	if isDue := utils.IsDue(db, "system", 10800); isDue {
		log.Printf("[INFO] MODULE: processor")

		processors, err := provider.Gather()
		if err != nil {
			log.Printf("[ERROR] %s", err)
			return
		}

		// Feeder
		if err := utils.Feeder(cfg, "/processor", processors); err != nil {
			log.Printf("[ERROR] %s", err)
			return
		}
	}

}
