package lpfs

import (
	//"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	DEFAULT_SAMPLES  int    = 1
	DEFAULT_INTERVAL int    = 1
	PROCDIR          string = "/proc"
	PROCDIR_LOADAVG  string = "/proc/loadavg"
	PROCDIR_MEMINFO  string = "/proc/meminfo"
	PROCDIR_SWAPS    string = "/proc/swaps"
	STATUSFILE       string = "stat"
)

//	GetLoadAverage1 returns the load average over the last minute.
func GetLoadAverage1() float64 {

	dat, err := os.ReadFile(PROCDIR_LOADAVG)
	if err != nil {
		// TODO
	}

	dat_s := strings.Split(string(dat), " ")[0]

	lavg1, err := strconv.ParseFloat(dat_s, 64)
	if err != nil {
		// TODO
	}

	return lavg1
}
