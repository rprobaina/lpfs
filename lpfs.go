package lpfs

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	procdir                  string = "/proc"
	procdir_loadavg          string = "/proc/loadavg"
	procdir_swaps            string = "/proc/swaps"
	procdir_stat             string = "/proc/stat"
	procdir_uptime           string = "/proc/uptime"
	procdir_per_process_stat string = "/stat"
	procdir_meminfo		 string = "/proc/meminfo"
	procdir_osrelease	 string = "/proc/sys/kernel/osrelease"
)

//	Procstat contains process stat available in /proc/<pid>/stat.
type Procstat struct {
	Pid                 int
	Comm                string
	State               string
	Ppid                int
	Pgrp                int
	Session             int
	TtyNr               int
	Tpgid               int
	Flags               int
	Minflt              int
	Cminflt             int
	Majflt              int
	Cmajflt             int
	Utime               int
	Stime               int
	Cutime              int
	Cstime              int
	Priority            int
	Nice                int
	NumThreads          int
	Itrealvalue         int
	Starttime           int
	Vsize               int
	Rss                 int
	Rsslim              string
	Startcode           int
	Endcode             int
	Startstack          int
	Kstkesp             int
	Kstkeip             int
	Signal              int
	Blocked             int
	Sigignore           int
	Sigcatch            int
	Wchan               int
	Nswap               int
	Cnswap              int
	ExitSignal          int
	Processor           int
	RtPriority          int
	Policy              int
	DelayacctBlkioTicks int
	GuestTime           int
	CguestTime          int
	StartData           int
	EndData             int
	StartBrk            int
	ArgStart            int
	ArgEnd              int
	EnvStart            int
	EnvEnd              int
	ExitCode            int
}

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
	if dat_s[1] == "" {
		return "", fmt.Errorf("no swap partition")
	}

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
	if dat_s[1] == "" {
		return "", fmt.Errorf("no swap partition")
	}

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
	if dat_s[1] == "" {
		return 0, fmt.Errorf("no swap partition")
	}

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
	if dat_s[1] == "" {
		return 0, fmt.Errorf("no swap partition")
	}

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
	if dat_s[1] == "" {
		return 0, fmt.Errorf("no swap partition")
	}

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

//	GetProcessesBlockedSize returns the number of blocked processes in the system.
// 	FIXME
func GetProcessesBlockedSize() (int, error) {
	dat, err := os.ReadFile(procdir_stat)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_stat)
		return 0, err
	}

	dat_s := strings.Split(string(dat), " ")[14]

	fmt.Println(dat_s)

	blksz, err := strconv.Atoi(strings.Split(dat_s, " ")[1])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return 0, err
	}

	return blksz, nil
}

//	GetPerProcessStat returns a slice of Procstat containing per-process (all living processes in the system) stat information.
func GetPerProcessStat() ([]Procstat, error) {

	var pps_s []Procstat

	files, err := ioutil.ReadDir(procdir)
	if err != nil {
		fmt.Printf("Error reading %s\n", procdir)
	}

	// Walking though /proc
	for _, f := range files {
		// FIXME: find a clever way of doing it.
		if f.Name() == "acpi" { // acpi is the first directory after the last PID.
			break
		}

		pid, err := strconv.Atoi(f.Name())
		if err != nil {
			fmt.Errorf("Error parsing %v to int\n", f.Name())
			return []Procstat{}, err
		}

		// get stats of <pid> via GetProcessStat()
		p, err := GetProcessStat(pid)
		if err != nil {
			fmt.Errorf("Error getting stats from PID %v\n", pid)
			return []Procstat{}, err
		}

		// Append p element into pps_s slice.
		pps_s = append(pps_s, p)
	}

	return pps_s, nil
}

//	GetProcessStat returns stat information of a giving process.
func GetProcessStat(pid int) (Procstat, error) {

	statFile := procdir + "/" + strconv.Itoa(pid) + "/" + procdir_per_process_stat

	dat, err := os.ReadFile(statFile)
	if err != nil {
		fmt.Errorf("Error reading %s\n", statFile)
		return Procstat{}, err
	}

	dat_s := strings.Split(string(dat), " ")

	var p Procstat

	// Parsing $statFile into Procstat fields.
	p.Pid, err = strconv.Atoi(dat_s[0])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Comm = string(dat_s[1])[1 : len(dat_s[1])-1]
	p.State = dat_s[2]

	p.Pgrp, err = strconv.Atoi(dat_s[4])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Session, err = strconv.Atoi(dat_s[5])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.TtyNr, err = strconv.Atoi(dat_s[6])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Tpgid, err = strconv.Atoi(dat_s[7])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Flags, err = strconv.Atoi(dat_s[8])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Minflt, err = strconv.Atoi(dat_s[9])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Cminflt, err = strconv.Atoi(dat_s[10])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Majflt, err = strconv.Atoi(dat_s[11])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Cmajflt, err = strconv.Atoi(dat_s[12])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Utime, err = strconv.Atoi(dat_s[13])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Stime, err = strconv.Atoi(dat_s[14])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Cutime, err = strconv.Atoi(dat_s[15])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Cstime, err = strconv.Atoi(dat_s[16])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Priority, err = strconv.Atoi(dat_s[17])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Nice, err = strconv.Atoi(dat_s[18])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.NumThreads, err = strconv.Atoi(dat_s[19])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Itrealvalue, err = strconv.Atoi(dat_s[20])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Starttime, err = strconv.Atoi(dat_s[21])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Vsize, err = strconv.Atoi(dat_s[22])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Rss, err = strconv.Atoi(dat_s[23])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Rsslim = dat_s[24]

	p.Startcode, err = strconv.Atoi(dat_s[25])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Endcode, err = strconv.Atoi(dat_s[26])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Startstack, err = strconv.Atoi(dat_s[27])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Kstkesp, err = strconv.Atoi(dat_s[28])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Kstkeip, err = strconv.Atoi(dat_s[29])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Signal, err = strconv.Atoi(dat_s[30])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Blocked, err = strconv.Atoi(dat_s[31])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Sigignore, err = strconv.Atoi(dat_s[32])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Sigcatch, err = strconv.Atoi(dat_s[33])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Wchan, err = strconv.Atoi(dat_s[34])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Nswap, err = strconv.Atoi(dat_s[35])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Cnswap, err = strconv.Atoi(dat_s[36])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.ExitSignal, err = strconv.Atoi(dat_s[37])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Processor, err = strconv.Atoi(dat_s[38])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.RtPriority, err = strconv.Atoi(dat_s[39])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.Policy, err = strconv.Atoi(dat_s[40])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.DelayacctBlkioTicks, err = strconv.Atoi(dat_s[41])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.GuestTime, err = strconv.Atoi(dat_s[42])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.CguestTime, err = strconv.Atoi(dat_s[43])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.StartData, err = strconv.Atoi(dat_s[44])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.EndData, err = strconv.Atoi(dat_s[45])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.StartBrk, err = strconv.Atoi(dat_s[46])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.ArgStart, err = strconv.Atoi(dat_s[47])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.ArgEnd, err = strconv.Atoi(dat_s[48])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.EnvStart, err = strconv.Atoi(dat_s[49])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.EnvEnd, err = strconv.Atoi(dat_s[50])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	p.ExitCode, err = strconv.Atoi(string(dat_s[51])[:len(dat_s[51])-1])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return Procstat{}, err
	}

	return p, nil
}

//	GetMemTotal returns the total memory
func GetMemTotal() (int, error) {
	dat, err := os.ReadFile(procdir_meminfo)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_meminfo)
	}

	dat_s := strings.Split(string(dat), "\n")

	s, err := strconv.Atoi(strings.Fields(dat_s[0])[1])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return 0, err
	}

	return s, err
}

//	GetMemFree returns the free memory
func GetMemFree() (int, error) {
	dat, err := os.ReadFile(procdir_meminfo)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_meminfo)
	}

	dat_s := strings.Split(string(dat), "\n")

	s, err := strconv.Atoi(strings.Fields(dat_s[1])[1])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return 0, err
	}

	return s, err
}

//	GetMemUsed returns the memory used
func GetMemUsed() (int, error) {
	dat, err := os.ReadFile(procdir_meminfo)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_meminfo)
	}

	dat_s := strings.Split(string(dat), "\n")

	t, err := strconv.Atoi(strings.Fields(dat_s[0])[1])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return 0, err
	}

	f, err := strconv.Atoi(strings.Fields(dat_s[1])[1])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return 0, err
	}

	return t - f, err
}

//	GetMemAvailable returns available memory
func GetMemAvailable() (int, error) {
	dat, err := os.ReadFile(procdir_meminfo)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_meminfo)
	}

	dat_s := strings.Split(string(dat), "\n")

	s, err := strconv.Atoi(strings.Fields(dat_s[2])[1])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return 0, err
	}

	return s, err
}

//	GetMemBuffers returns the memory buffers
func GetMemBuffers() (int, error) {
	dat, err := os.ReadFile(procdir_meminfo)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_meminfo)
	}

	dat_s := strings.Split(string(dat), "\n")

	s, err := strconv.Atoi(strings.Fields(dat_s[3])[1])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return 0, err
	}

	return s, err
}

//	GetMemCached returns the memory cached
func GetMemCached() (int, error) {
	dat, err := os.ReadFile(procdir_meminfo)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_meminfo)
	}

	dat_s := strings.Split(string(dat), "\n")

	s, err := strconv.Atoi(strings.Fields(dat_s[4])[1])
	if err != nil {
		fmt.Errorf("error parsing %v", dat_s)
		return 0, err
	}

	return s, err
}

//	GetKernelRelease returns the kernel version with additional information.
func GetKernelRelease() (string, error) {
	dat, err := os.ReadFile(procdir_osrelease)
	if err != nil {
		fmt.Errorf("unable to read the file %v", procdir_osrelease)
		return "", err
	}

	return string(dat), err
}
