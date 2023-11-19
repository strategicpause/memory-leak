package memory

import (
	"crypto/rand"
	"fmt"
	"github.com/strategicpause/memory-leak/metrics"
	"time"
)

type Params struct {
	MaxMemoryInBytes   uint64
	BlockSizeInBytes   uint64
	StepTimeInSeconds  time.Duration
	PauseTimeInSeconds time.Duration
	RandomData         bool
}

func memoryLeak(params *Params) error {
	PrintParams(params)

	numEntries := int(params.MaxMemoryInBytes / params.BlockSizeInBytes)
	list := make([][]byte, numEntries)

	for i := 0; i < numEntries; i++ {
		list[i] = make([]byte, params.BlockSizeInBytes)
		if params.RandomData {
			_, _ = rand.Read(list[i])
		} else {
			// This will result in allocating memory from the virtual address space.
			for j := 0; j < int(params.BlockSizeInBytes); j++ {
				list[i][j] = 0
			}
		}
		metrics.PrintMemory()
		time.Sleep(params.StepTimeInSeconds)
	}

	fmt.Printf("Waiting for %s.\n", params.PauseTimeInSeconds.String())
	time.Sleep(params.PauseTimeInSeconds)
	fmt.Println("Done")

	return nil
}

func PrintParams(params *Params) {
	fmt.Printf("MaxMemory = %v MiB\tBlockSize = %v MiB\tPauseTime = %v\tRandomData = %t.\n",
		metrics.BToMiB(params.MaxMemoryInBytes), metrics.BToMiB(params.BlockSizeInBytes),
		params.PauseTimeInSeconds, params.RandomData)
}
