package memory

import (
	"errors"
	"github.com/inhies/go-bytesize"
	"github.com/urfave/cli"
)

const (
	MaxMemoryName      = "max-memory"
	BlockSizeName      = "block-size"
	StepDurationName   = "step-duration"
	PauseDurationName  = "pause-duration"
	WithRandomDataName = "with-random-data"
)

func Register() cli.Command {
	return cli.Command{
		Name:   "memory",
		Usage:  "Reproduces a memory leak.",
		Action: action,
		Flags:  flags(),
	}
}

func flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  MaxMemoryName,
			Usage: "Specify the maximum amount of memory to acquire.",
			Value: "1GB",
		},
		cli.StringFlag{
			Name:  BlockSizeName,
			Usage: "Specify the block size of memory which will be allocated at any given time.",
			Value: "10MB",
		},
		cli.DurationFlag{
			Name:  StepDurationName,
			Usage: "Time between allocations in seconds.",
			Value: 1,
		},
		cli.DurationFlag{
			Name:  PauseDurationName,
			Usage: "Time to wait, in seconds, after all memory has been allocated.",
			Value: 0,
		},
		cli.BoolFlag{
			Name: WithRandomDataName,
		},
	}
}

func action(ctx *cli.Context) error {
	maxMemory, err := bytesize.Parse(ctx.String(MaxMemoryName))
	if err != nil {
		return err
	}

	blockSize, err := bytesize.Parse(ctx.String(BlockSizeName))
	if err != nil {
		return err
	}

	if blockSize > maxMemory {
		return errors.New("block-size must be less than or equal to max-memory")
	}

	params := &Params{
		MaxMemoryInBytes:   uint64(maxMemory),
		BlockSizeInBytes:   uint64(blockSize),
		StepTimeInSeconds:  ctx.Duration(StepDurationName),
		PauseTimeInSeconds: ctx.Duration(PauseDurationName),
		RandomData:         ctx.Bool(WithRandomDataName),
	}

	return memoryLeak(params)
}
