package workers

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var sample = []string{
	"hello",
	"are you there?",
	"okay",
	"i am busy",
}

func Worker1Pub(out chan<- string, wg *sync.WaitGroup, sleepOffset int, workerNumber int) {
	for i := 0; i < 2; i++ {
		for _, k := range sample {
			out <- k
			fmt.Println("sent: " + k + " \"by worker number: " + strconv.Itoa(workerNumber) + "\"")
			time.Sleep(time.Second * time.Duration(sleepOffset))
		}
		time.Sleep(time.Second * time.Duration(sleepOffset))
	}
	wg.Done()
}
