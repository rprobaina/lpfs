// This is a simple example of lpfs usage that simulates the output of `$sar -q`.
package main

import (
	"fmt"
	"github.com/rprobaina/lpfs"
	"time"
)

func main() {
	fmt.Println("\t\trunq\tplist\t lavg1\tlavg5\tlavg15\tblkd")
	for {
		runq, _ := lpfs.GetRunnableQueueSize()
		plist, _ := lpfs.GetTaskQueueSize()
		lavg1, _ := lpfs.GetLoadAverage1()
		lavg5, _ := lpfs.GetLoadAverage5()
		lavg15, _ := lpfs.GetLoadAverage15()
		blk := 0 // FIXME
		t := time.Now().Format("15:04:05")
		fmt.Printf("%v\t%v\t%v\t %.1f\t%.1f\t%.1f\t%v\n", t, runq, plist, lavg1, lavg5, lavg15, blk)
		time.Sleep(time.Second)
	}
}
