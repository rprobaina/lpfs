package lpfs

import (
	//"io/ioutil"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	procdir         string = "/proc"
	procdir_loadavg string = "/proc/loadavg"
	procdir_swaps   string = "/proc/swaps"
	procdir_uptime  string = "/proc/uptime"
)

//	GetLoadAverage1 returns the load average over the last minute.
func GetLoadAverage1() (float64, error) {
	dat, err := os.ReadFile(procdir_loadavg)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_loadavg)
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
	dat, err := os.ReadFile(procdir_loadavg)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_loadavg)
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
	dat, err := os.ReadFile(procdir_loadavg)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_loadavg)
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
	dat, err := os.ReadFile(procdir_loadavg)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_loadavg)
		return 0, err
	}

	dat_s := strings.Split(string(dat), " ")[3]

	runq, err := strconv.Atoi(strings.Split(dat_s, "/")[0])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return 0, err
	}

	return runq, err
}

//	GetTaskQueueSize returns the number of existing tasks in the system.
func GetTaskQueueSize() (int, error) {
	dat, err := os.ReadFile(procdir_loadavg)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_loadavg)
		return 0, err
	}

	dat_s := strings.Split(string(dat), " ")[3]

	tskq, err := strconv.Atoi(strings.Split(dat_s, "/")[1])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return 0, err
	}

	return tskq, err
}

//	GetMostRecentPid returns the the PID of the process that was most recently created on the system.
func GetMostRecentPid() (int, error) {
	dat, err := os.ReadFile(procdir_loadavg)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_loadavg)
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
	dat, err := os.ReadFile(procdir_swaps)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_swaps)
		return "", err
	}

	dat_s := strings.Split(string(dat), "\n")

	s := strings.Split(strings.Join(strings.Fields(strings.TrimSpace(dat_s[1])), " "), " ")[0]

	return s, err
}

//	GetSwapType returns the swap partition type.
func GetSwapType() (string, error) {
	dat, err := os.ReadFile(procdir_swaps)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_swaps)
		return "", err
	}

	dat_s := strings.Split(string(dat), "\n")

	s := strings.Split(strings.Join(strings.Fields(strings.TrimSpace(dat_s[1])), " "), " ")[1]

	return s, err
}

//	GetSwapSize returns the swap partition total size.
func GetSwapSize() (int, error) {
	dat, err := os.ReadFile(procdir_swaps)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_swaps)
		return 0, err
	}

	dat_s := strings.Split(string(dat), "\n")

	s, err := strconv.Atoi(strings.Split(strings.Join(strings.Fields(strings.TrimSpace(dat_s[1])), " "), " ")[2])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return 0, err
	}

	return s, err
}

//	GetSwapUsed returns the swap partition used size.
func GetSwapUsed() (int, error) {
	dat, err := os.ReadFile(procdir_swaps)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_swaps)
		return 0, err
	}

	dat_s := strings.Split(string(dat), "\n")

	s, err := strconv.Atoi(strings.Split(strings.Join(strings.Fields(strings.TrimSpace(dat_s[1])), " "), " ")[3])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return 0, err
	}

	return s, err
}

//	GetSwapPriority returns the swap partition priority.
func GetSwapPriority() (int, error) {
	dat, err := os.ReadFile(procdir_swaps)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_swaps)
		return 0, err
	}

	dat_s := strings.Split(string(dat), "\n")

	s, err := strconv.Atoi(strings.Split(strings.Join(strings.Fields(strings.TrimSpace(dat_s[1])), " "), " ")[4])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return 0, err
	}

	return s, err
}

//	GetUptimeSystem returns the uptime of the system (seconds).
func GetUptimeSystem() (float64, error) {
	dat, err := os.ReadFile(procdir_uptime)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_uptime)
		return 0.0, err
	}

	dat_s := strings.Split(string(dat), " ")[0]

	us, err := strconv.ParseFloat(dat_s, 32)
	if err != nil {
		fmt.Errorf("error parsing %v to float", dat_s)
		return 0.0, err
	}

	return us, nil
}

//	GetUptimeIdle returns the amount of time spent in idle process (seconds).
func GetUptimeIdle() (float64, error) {
	dat, err := os.ReadFile(procdir_uptime)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_uptime)
		return 0.0, err
	}

	dat_s := strings.Split(string(dat), " ")[1]

	ui, err := strconv.ParseFloat(dat_s[:len(dat_s)-1], 32)
	if err != nil {
		fmt.Errorf("error parsing %v to float", dat_s)
		return 0.0, err
	}

	return ui, nil
}
