package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("start")
	// uncomment which ones to test
	// loopDefer() // defers are stacked and performed LIFO after function scope has finished
	// loopRoutine() // go routines starts another worker thread that may be preempted when main thread exist (no waiting)
	// loopRoutineWithSleep() // go routine with sleep gives enough time for the worker thread to perform
	loopRoutineWithWait() // similar to routine with sleep but with more graceful approach to signalling and waiting

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("outer", i)
	// }

	fmt.Println("end")
}

func loopDefer() {
	fmt.Println("counting with defer...")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func loopRoutine() {
	fmt.Println("counting with go routines...")
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}

	fmt.Println("done")
}

func loopRoutineWithSleep() {
	fmt.Println("counting with go routines and sleep...")
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}

	time.Sleep(time.Millisecond)

	fmt.Println("done")
}

func loopRoutineWithWait() {
	fmt.Println("counting with go routines and wait group...")

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			fmt.Println(i)
			defer wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("done")
}
