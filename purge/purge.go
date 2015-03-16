package purge

import (
	"encoding/json"
	"github.com/op/go-logging"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

var log = logging.MustGetLogger("purge")
var format = logging.MustStringFormatter(
	"%{color}%{time:2006-01-02 15:04:05} %{pid} %{level:.6s} %{id:03x}%{color:reset} %{message}",
)

type PurgeConf struct {
	PurgeInfos     []PurgeInfo `json:"purge_info"`
	ParallelDegree int         `json:"parallel_degree"`
}

type PurgeInfo struct {
	Path  string `json:"path"`
	Delay int    `json:"delay"`
}

// Wait group to control parallel execution
var wg sync.WaitGroup

// Run the purge, using the given configuration file.
// Do purge of listed couples paths/delay in parallel.
func Run(filename string) int {
	InitLog()

	log.Info("Start purge")
	file, e := ioutil.ReadFile(filename)

	if e != nil {
		log.Error("File error: %v\n", e)
		os.Exit(1)
	}

	var purgeConf PurgeConf
	var purgeInfos []PurgeInfo

	err := json.Unmarshal(file, &purgeConf)
	if err != nil {
		log.Error("Error when parsing json file %s", err)
		return 1
	}

	runtime.GOMAXPROCS(purgeConf.ParallelDegree)
	purgeInfos = purgeConf.PurgeInfos

	for i := 0; i < len(purgeInfos); i++ {
		wg.Add(1)
		go Purge(purgeInfos[i])
	}

	log.Info("Waiting parallel purge to finish...")
	wg.Wait()
	log.Info("Purge done.")
	return 0
}

// Init the log for purge.
func InitLog() {
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
}

// Purge one path when files inside are older than configured delay
func Purge(info PurgeInfo) {
	defer wg.Done()

	log.Info("Process line: path=%s, delay=%d", info.Path, info.Delay)

	if info.Path == "" {
		log.Error("Path is undefined")
	} else {
		files, _ := filepath.Glob(info.Path)

		for i := 0; i < len(files); i++ {
			filePath := files[i]

			if isPurgeable(info.Delay, filePath) {
				e := os.Remove(filePath)
				if e != nil {
					log.Error("Error when deleting file %s", filePath)
				} else {
					log.Info("File %s deleted successfully", filePath)
				}
			}
		}
	}
}

// Return true if the given file has modification time older than (current time - delay).
func isPurgeable(delay int, filePath string) bool {
	timeLimit := time.Now().Add(-(time.Duration(delay*24) * time.Hour))

	fi, err := os.Stat(filePath)
	return err != nil || (timeLimit.After(fi.ModTime()) && !fi.IsDir())
}
