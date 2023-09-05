package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

const (
	MiB     = 1024 * 1024
	Port    = ":8080"
	TcpLeak = true
)

func main() {
	if TcpLeak {
		tcpLeak()
	} else {
		memoryLeak()
	}

}

func tcpLeak() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("Listening on port", Port)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(time.Now().String()))
		})
		http.ListenAndServe(Port, nil)
		wg.Done()
	}()
	for {
		go func() {
			req := Must(http.NewRequest("GET", "http://localhost:8080/", nil))
			client := http.Client{}
			Must(client.Do(req))
			PrintMemory()
		}()
	}
	wg.Wait()
}

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}
	return obj
}

func memoryLeak() {
	// Allocate 10 GiB in 100 MiB increments
	list := make([][]byte, 100)
	for i := 0; i < 100; i++ {
		list[i] = make([]byte, 100*MiB)
		for j := 0; j < 100*MiB; j++ {
			list[i][j] = 0
		}
		PrintMemory()
		time.Sleep(time.Second)
	}
	fmt.Println("Done")
}

func PrintMemory() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc = %v MiB\tSys = %v MiB\n", bToMiB(m.TotalAlloc), bToMiB(m.Sys))
}

func bToMiB(b uint64) uint64 {
	return b / 1024 / 1024
}
