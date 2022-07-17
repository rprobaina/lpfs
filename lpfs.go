package lpfs

import (
	//"io/ioutil"
	"fmt"
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
		fmt.Errorf("unable to read the file %v", PROCDIR_LOADAVG)
		return 0.0, err
	}

	dat_s := strings.Split(string(dat), " ")[0]

	lavg, err := strconv.ParseFloat(dat_s, 32)
	if err != nil {
		fmt.Errorf("error parsing %v to float", dat_s)
		return 0.0, err
	}

	return lavg, nil
}

//	GetLoadAverage5 returns the load average over the last 5 minutes.
func GetLoadAverage5() (float64, error) {
	dat, err := os.ReadFile(PROCDIR_LOADAVG)
	if err != nil {
		fmt.Errorf("unable to read the file %v", PROCDIR_LOADAVG)
		return 0.0, err
	}

	dat_s := strings.Split(string(dat), " ")[1]

	lavg, err := strconv.ParseFloat(dat_s, 32)
	if err != nil {
		fmt.Errorf("error parsing %v to float", dat_s)
		return 0.0, err
	}

	return lavg, nil
}

//	GetLoadAverage15 returns the load average over the last 15 minutes.
func GetLoadAverage15() (float64, error) {
	dat, err := os.ReadFile(PROCDIR_LOADAVG)
	if err != nil {
		fmt.Errorf("unable to read the file %v", PROCDIR_LOADAVG)
		return 0.0, err
	}

	dat_s := strings.Split(string(dat), " ")[2]

	lavg, err := strconv.ParseFloat(dat_s, 32)
	if err != nil {
		fmt.Errorf("error parsing %v to float", dat_s)
		return 0.0, err
	}

	return lavg, nil
}


//	GetRunnableQueueSize returns the number of currently runnable tasks.
func GetRunnableQueueSize() (int, error) {
	dat, err := os.ReadFile(PROCDIR_LOADAVG)
	if err != nil {
		fmt.Errorf("unable to read the file %v", PROCDIR_LOADAVG)
		return 0, err
	}

	dat_s := strings.Split(string(dat), " ")[3]

	runq, err := strconv.Atoi(strings.Split(dat_s, "/")[0])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return 0, err
	}

	return  runq, err
}

//	GetTaskQueueSize returns the number of existing tasks in the system.
func GetTaskQueueSize() (int, error) {
	dat, err := os.ReadFile(PROCDIR_LOADAVG)
	if err != nil {
		fmt.Errorf("unable to read the file %v", PROCDIR_LOADAVG)
		return 0, err
	}

	dat_s := strings.Split(string(dat), " ")[3]

	tskq, err := strconv.Atoi(strings.Split(dat_s, "/")[1])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return 0, err
	}

	return  tskq, err
}

//	GetMostRecentPid returns the the PID of the process that was most recently created on the system.
func GetMostRecentPid() (int, error) {
	dat, err := os.ReadFile(PROCDIR_LOADAVG)
	if err != nil {
		fmt.Errorf("unable to read the file %v", PROCDIR_LOADAVG)
		return 0, err
	}

	dat_s := strings.Split(string(dat), " ")[4]

	pid, err := strconv.Atoi(dat_s[:len(dat_s)-1])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return 0, err
	}

	return pid, nil
}


//	GetSwapFilename returns the swap partition filename.
func GetSwapFilename() (string, error) {
        dat, err := os.ReadFile(PROCDIR_SWAPS)
        if err != nil {
		fmt.Errorf("unable to read the file %v", PROCDIR_SWAPS)
		return "", err
        }

        dat_s := strings.Split(string(dat), "\n")

        s := strings.Split(strings.Join(strings.Fields(strings.TrimSpace(dat_s[1])), " "), " ")[0]

        return s, err;
}
