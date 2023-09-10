package metrics

import (
	"fmt"
	"runtime"
)

func PrintMemory() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	var avgGCTime uint64
	if m.NumGC > 0 {
		avgGCTime = NsToUs(m.PauseTotalNs / uint64(m.NumGC))
	}

	fmt.Printf("TotalAlloc = %v MiB\tHeapAlloc = %v MiB\tSys = %v MiB\tNextGC = %v MiB\tPauseTotalUs = %v\tNumGC = %v\tAvgGCTime = %v\n",
		// Total number of bytes allocated for the heap. Does not decrease when objects are freed.
		BToMiB(m.TotalAlloc),
		// Total number of bytes allocated on the heap. This includes reachable objects and objects
		// not yet freed by the GC. This value should decrease when objects are cleaned by the GC.
		BToMiB(m.HeapAlloc),
		// The total number of bytes obtained from the OS.
		BToMiB(m.Sys),
		// The target heap size of the next GC cycle. The GC goal is to kep HeapAlloc <= NextGC.
		BToMiB(m.NextGC),
		// The cumulative nanoseconds in GC stop-the-world pauses since the program started.
		NsToUs(m.PauseTotalNs),
		// Number of completed GC cycles.
		m.NumGC,
		// Average time spent in GC across. Pretty hacky way to get some idea of how GC time is being spent during the leak.
		avgGCTime)
}

func BToMiB[T uint64 | uint32](b T) uint64 {
	return uint64(b) / (1024 * 1024)
}

func NsToUs(s uint64) uint64 {
	return s / 1000
}
