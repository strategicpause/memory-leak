package tcp

import (
	"fmt"
	"github.com/strategicpause/memory-leak/metrics"
	"net/http"
	"time"
)

func tcpLeak() error {
	go func() {
		fmt.Println("Listening on port", Port)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(time.Now().String()))
		})
		_ = http.ListenAndServe(Port, nil)
	}()
	for {
		go func() {
			req := Must(http.NewRequest("GET", "http://localhost:8080/", nil))
			client := http.Client{}
			Must(client.Do(req))
			metrics.PrintMemory()
		}()
	}
}

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}
	return obj
}
