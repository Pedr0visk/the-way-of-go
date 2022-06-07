package main

import (
	"runtime"
	"time"
)

func main() {
	ms := runtime.MemStats{}
	runtime.ReadMemStats(&ms)

	println("Heap after GC. Used in Kb:", ms.HeapInuse, " Free:", ms.HeapIdle, " Meta:", ms.GCSys)

	time.Sleep(5 * time.Second)
}
