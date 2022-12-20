package profiler

import (
	"fmt"
	fact "msg-app/api/factory"
	"runtime"
	"time"
)

func MemoryProfiler() {
	done := make(chan bool)
	fileWriter := fact.FileSenderFactory()
	fileToWrite := "profiler/memoria.txt"
	var lastAlloc uint64
	var mem runtime.MemStats
	var memUsage uint64

	go func() {
		for {
			currentTime := time.Now().Format("01-02-2006 15:04:05")
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
