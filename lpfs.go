package lpfs

import (
	//"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	PROCDIR         string = "/proc"
	PROCDIR_LOADAVG string = "/proc/loadavg"
	PROCDIR_MEMINFO string = "/proc/meminfo"
	PROCDIR_SWAPS   string = "/proc/swaps"
	STATUSFILE      string = "stat"
)

//	GetLoadAverage1 returns the load average over the last minute.
func GetLoadAverage1() (float64, error) {
	dat, err := os.ReadFile(PROCDIR_LOADAVG)
	if err != nil {
		return 0.0, err
	}

	dat_s := strings.Split(string(dat), " ")[0]

	lavg, err := strconv.ParseFloat(dat_s, 32)
	if err != nil {
		return 0.0, err
	}

	return lavg, nil
}

//	GetLoadAverage5 returns the load average over the last 5 minutes.
func GetLoadAverage5() (float64, error) {
	dat, err := os.ReadFile(PROCDIR_LOADAVG)
	if err != nil {
		return 0.0, err
	}

	dat_s := strings.Split(string(dat), " ")[1]

	lavg, err := strconv.ParseFloat(dat_s, 32)
	if err != nil {
		return 0.0, err
	}

	return lavg, nil
}

//	GetLoadAverage15 returns the load average over the last 15 minutes.
func GetLoadAverage15() (float64, error) {
	dat, err := os.ReadFile(PROCDIR_LOADAVG)
	if err != nil {
		return 0.0, err
	}

	dat_s := strings.Split(string(dat), " ")[2]

	lavg, err := strconv.ParseFloat(dat_s, 32)
	if err != nil {
		return 0.0, err
	}

	return lavg, nil
}
