// This is a example of functions that get per-process information.
// This particular program shows the state of all processes in the system.
package main

import (
	"fmt"
	"github.com/rprobaina/lpfs"
)

func main() {
	fmt.Println("STAT\tCOMM (PID)")
	pps, err := lpfs.GetPerProcessStat()
	if err != nil {
		fmt.Errorf("Unable to get per-process stat info. Error: %v", err)
	}

	for _, i := range pps {
		fmt.Printf("%v\t%v (%v)\n", i.State, i.Comm, i.Pid)
	}
}
