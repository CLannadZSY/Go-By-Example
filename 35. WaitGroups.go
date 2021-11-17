package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func workerWaitGroups(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func workerRun(id int) {
	workerWaitGroups(id)
	wg.Done()
}

func main() {

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		//这种写法, 即使不使用闭包, 也不会导致重复使用相同的值
		//i := i

		go workerRun(i)
	}

	wg.Wait()

}
