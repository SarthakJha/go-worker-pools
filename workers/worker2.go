package workers

import (
	"fmt"
	"strconv"
	"sync"
)

func Worker2Sub(in <-chan string, wg *sync.WaitGroup, workerNumber int) {
	for val := range in {
		fmt.Println("						" + "recieved: " + val + " worker number: " + strconv.Itoa(workerNumber))
	}
}
