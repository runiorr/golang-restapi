package profiler

import (
	"fmt"
	fact "msg-app/api/factory"
	"runtime"
	"time"
)

func MemoryProfiler() {
	done := make(chan bool)
	fileToWrite := "memoria.txt"
	fileWriter := fact.FileSenderFactory()

	go func() {
		var lastAlloc uint64
		for {
			currentTime := time.Now().Format("01-02-2006 15:04:05")

			var mem runtime.MemStats
			var memUsage uint64
			runtime.ReadMemStats(&mem)
			if lastAlloc > 0 {
				memUsage = mem.Alloc - lastAlloc
			}
			lastAlloc = mem.Alloc
			dataPoint := fmt.Sprintf("datetime: %s, mem_used(bytes): %d\n", currentTime, memUsage)
			fileWriter.Send(fileToWrite, dataPoint)
			done <- true
		}
	}()

	for {
		<-done
		time.Sleep(time.Second)
	}
}
