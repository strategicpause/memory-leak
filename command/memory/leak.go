package memory

import (
	"arena"
	"fmt"
	"github.com/strategicpause/memory-leak/metrics"
	"time"
)

type Params struct {
	MaxMemoryInBytes   uint64
	BlockSizeInBytes   uint64
	StepTimeInSeconds  time.Duration
	PauseTimeInSeconds time.Duration
}

func memoryLeak(params *Params) error {

	PrintParams(params)

	entries := int(params.MaxMemoryInBytes / (1024 * 1024))

	a := arena.NewArena()
	defer a.Free()

	slice := make([][]byte, entries)

	for i := 0; i < entries; i++ {
		slice[i] = arena.MakeSlice[byte](a, 1024*1024, 1024*1024)
		metrics.PrintMemory()
	}

	fmt.Printf("Waiting for %s.\n", params.PauseTimeInSeconds.String())
	time.Sleep(params.PauseTimeInSeconds)
	fmt.Println("Done")

	return nil
}

func PrintParams(params *Params) {
	fmt.Printf("MaxMemory = %v MiB\tBlockSize = %v MiB\tPauseTime = %v.\n",
		metrics.BToMiB(params.MaxMemoryInBytes), metrics.BToMiB(params.BlockSizeInBytes), params.PauseTimeInSeconds)
}
