package cluster

import (
	"log"
	"runtime"
	"strconv"
	"strings"
	"tgo/api/internal/config"
)

func SetMaxProcs() {
	maxProcsEnv := strings.ToUpper(config.GetEnv("MAX_PROCS", "MAX"))

	numCPU := runtime.NumCPU()

	if maxProcsEnv == "MAX" {
		runtime.GOMAXPROCS(numCPU)
		log.Printf("[ClusterModule] Utilizando todos os %d núcleos disponíveis.", numCPU)
	} else {

		maxProcs, err := strconv.Atoi(maxProcsEnv)
		if err != nil || maxProcs <= 0 {
			maxProcs = 1
		}
		runtime.GOMAXPROCS(maxProcs)
		log.Printf("[ClusterModule] Utilizando %d núcleo(s) de CPU.", maxProcs)
	}
}
