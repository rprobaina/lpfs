package lpfs

import (
	"fmt"
	"testing"
)

//	TestLoadAverage test all functions that get data from /proc/loadavg.
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
