package main

import (
	"fmt"
	"sync"

	"github.com/SarthakJha/work-pool/workers"
)

/*
	if this is put in main, deadlock will be faced
	ALWAYS DECLARE THIS GLOBALLY
*/
var wg sync.WaitGroup

func main() {
	// // unbuffered
	// strChan := make(chan string)
	strChan := make(chan string, 3)

	// creating worker pool of 10 workers
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go workers.Worker1Pub(strChan, &wg, i+1, i+1)
	}
	go workers.Worker2Sub(strChan, &wg, 1)
	go workers.Worker2Sub(strChan, &wg, 2)

	defer close(strChan)
	wg.Wait()
	fmt.Println("FIN")
}
