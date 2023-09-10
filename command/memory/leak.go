package memory

import (
	"fmt"
	"github.com/strategicpause/memory-leak/metrics"
	"time"
)

type Params struct {
	MaxMemoryInBytes   uint64
	BlockSizeInBytes   uint64
	PauseTimeInSeconds time.Duration
}

func memoryLeak(params *Params) error {
	PrintParams(params)

	numEntries := int(params.MaxMemoryInBytes / params.BlockSizeInBytes)
	list := make([][]byte, numEntries)

	for i := 0; i < numEntries; i++ {
		list[i] = make([]byte, params.BlockSizeInBytes)
		for j := 0; j < int(params.BlockSizeInBytes); j++ {
			list[i][j] = 0
		}
		metrics.PrintMemory()
		time.Sleep(params.PauseTimeInSeconds)
	}

	fmt.Println("Done")

	return nil
}

func PrintParams(params *Params) {
	fmt.Printf("MaxMemory = %v MiB\tBlockSize = %v MiB\tPauseTime = %v.\n",
		metrics.BToMiB(params.MaxMemoryInBytes), metrics.BToMiB(params.BlockSizeInBytes), params.PauseTimeInSeconds)
}
