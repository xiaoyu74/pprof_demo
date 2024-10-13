package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func heavyComputation() {
	for i := 0; i < 1e7; i++ {
		_ = rand.Float64() * rand.Float64()
	}
}

func memoryLeak() {
	data := make([][]byte, 0)
	for {
		// Allocate 1MB
		chunk := make([]byte, 1024*1024)
		// Hold the memory to mock mem leak
		data = append(data, chunk)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go func() {
		for {
			heavyComputation()
			// Mock CPU spikes
			time.Sleep(2 * time.Second)
		}
	}()

	go memoryLeak()

	fmt.Println("Starting pprof server on http://localhost:6060/debug/pprof/")
	http.ListenAndServe("localhost:6060", nil)
}
