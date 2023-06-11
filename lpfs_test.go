package lpfs

import (
	"fmt"
	"testing"
)

//	TestLoadAverage tests all functions that get data from /proc/loadavg.
func TestLoadAverage(t *testing.T) {
	l1, err := GetLoadAverage1()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("LoadAverage1(): %v, err: %v\n", l1, err)

	l5, err := GetLoadAverage5()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("LoadAverage5(): %v, err: %v\n", l5, err)

	l15, err := GetLoadAverage15()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("LoadAverage15(): %v, err: %v\n", l15, err)

	tskq, err := GetTaskQueueSize()
	if err != nil {
		t.Errorf("%v", err)
	}

	fmt.Printf("GetTaskQueueSize(): %v, err: %v\n", tskq, err)
	rszq, err := GetRunnableQueueSize()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetRunnableQueueSize(): %v, err: %v\n", rszq, err)

	pid, err := GetMostRecentPid()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetMostRecentPid(): %v, err: %v\n", pid, err)
}

//	TestSwaps tests all functions that get data from /proc/swaps.
func TestSwaps(t *testing.T) {
	sf, err := GetSwapFilename()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetSwapFilename(): %v, err: %v\n", sf, err)

	st, err := GetSwapType()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetSwapType(): %v, err: %v\n", st, err)

	ss, err := GetSwapSize()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetSwapSize(): %v, err: %v\n", ss, err)

	su, err := GetSwapUsed()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetSwapUsed(): %v, err: %v\n", su, err)

	sp, err := GetSwapPriority()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetSwapPriority(): %v, err: %v\n", sp, err)
}

//	TestUptime tests all functions that get data from /proc/uptime.
func TestUptime(t *testing.T) {
	us, err := GetUptimeSystem()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetUptimeSystem(): %v, err: %v\n", us, err)

	ui, err := GetUptimeIdle()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetUptimeIdle(): %v, err: %v\n", ui, err)
}

//	TestStat tests all functions that get data from /proc/stat.
/*
func TestStat(t *testing.T) {
	blksz, err := GetProcessesBlockedSize()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetProcessesBlockedSize(): %v, err: %v\n", blksz, err)
}
*/

// 	TestPerProcess tests all functions that get data from /proc/<pid>/.
func TestPerProcess(t *testing.T) {
	ppStat, err := GetPerProcessStat()
	if err != nil {
		t.Errorf("%#v", err)
	}
	fmt.Printf("GetPerProcessStat(): %v, err: %v\n", ppStat, err)

	ps, err := GetProcessStat(1)	
	if err != nil {
		t.Errorf("%#v", err)
	}
	fmt.Printf("GetProcessStat(1): %v, err: %v\n", ps, err)
}

//	TestMeminfo tests all functions that get from /proc/meminfo
func TestMeminfo(t *testing.T) {

	mt, err := GetMemTotal()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetMemTotal(): %v, err: %v\n", mt, err)

	mf, err := GetMemFree()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetMemFree(): %v, err: %v\n", mf, err)

	mu, err := GetMemUsed()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetMemUsed(): %v, err: %v\n", mu, err)

	ma, err := GetMemAvailable()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetMemAvailable(): %v, err: %v\n", ma, err)

	mb, err := GetMemBuffers()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetMemBuffers(): %v, err: %v\n", mb, err)

	mc, err := GetMemCached()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetMemCached(): %v, err: %v\n", mc, err)
}

//	TestSysKernel tests all functions that get data from /proc/sys/kernel.
func TestSysKernel(t *testing.T) {
	kr , err := GetKernelRelease()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetKernelRelease(): %v, err: %v\n", kr, err)
}
