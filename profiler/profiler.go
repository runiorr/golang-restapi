package profiler

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func MemoryProfiler() {
	done := make(chan bool)

	go func() {
		var lastAlloc uint64
		for {

			f, err := os.OpenFile("memoria.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()

			currentTime := time.Now().Format("01-02-2006 15:04:05")

			var mem runtime.MemStats
			runtime.ReadMemStats(&mem)

			var memUsage uint64
			if lastAlloc > 0 {
				memUsage = mem.Alloc - lastAlloc
			}
			lastAlloc = mem.Alloc

			if _, err := f.WriteString(
				fmt.Sprintf("datetime: %s, mem_used(bytes): %d\n",
					currentTime, memUsage)); err != nil {
				fmt.Println(err)
				return
			}

			done <- true
		}
	}()

	for {
		<-done
		time.Sleep(time.Second)
	}
}
