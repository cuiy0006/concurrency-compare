package main

import (
	"fmt"
	"net/http"
	"time"
)

//Go async average time for  1000  HEAD requests
//638.133765ms

const (
	IterationCount = 10
	RequestCount = 1000
	Url = "http://example.com"
)

func main() {
	var totalDuration time.Duration
	for i := 0; i < IterationCount; i++{
		totalDuration += testHead()
	}
	fmt.Println("Go async average time for ", RequestCount, " HEAD requests")
	fmt.Println(totalDuration / IterationCount)
}

func testHead() time.Duration{
	start := time.Now()
	revChannel, errChannel := make(chan bool), make(chan error)
	// HEAD request RequestCount times
	for i := 0; i < RequestCount; i++ {
		go func() {
			_, err := http.Head(Url)
			if err != nil {
				errChannel <- err
				return
			} else {
				revChannel <- true
				return
			}
		}()
	}
	//receive response
	for i := 0; i < 1000; i++{
		select {
		case <- revChannel:
			continue
		case err := <- errChannel:
			fmt.Println(err)
		}
	}

	return time.Since(start)
}
