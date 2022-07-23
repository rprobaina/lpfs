// This is a example of functions that get per-process information.
// This particular program shows the top Resident Set Size (RSS) memory consumers in the system.
package main

import (
	"fmt"
	"github.com/rprobaina/lpfs"
	"sort"
)

func main() {
	fmt.Println("RSS (MiB)\tCOMM (PID)")
	pps, err := lpfs.GetPerProcessStat()
	if err != nil {
		fmt.Errorf("Unable to get per-process stat info. Error: %v", err)
	}

	// Sorting pps slice by top RSS consumers.
	sort.Slice(pps, func(i, j int) bool {
		return pps[i].Rss > pps[j].Rss
	})

	for _, i := range pps {
		fmt.Printf("%v\t\t%v (%v)\n", i.Rss, i.Comm, i.Pid)
	}
}
