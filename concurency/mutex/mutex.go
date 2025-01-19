package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	counter := 0
	
	runCnt := 20
	wg.Add(runCnt)
	var mu sync.Mutex

	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	for i := 0; i < runCnt; i++ {
		go func() {
			mu.Lock()
			counter += 1
			//runtime.Gosched()
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(counter)
}
