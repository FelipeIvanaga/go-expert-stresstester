package httptester

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func Execute(url string, requests, concurrency int) {
	start := time.Now()
	var wg sync.WaitGroup
	ch := make(chan interface{}, concurrency)
	chStatus := make(chan int, requests)

	for i := 0; i < requests; i++ {
		wg.Add(1)
		ch <- 1
		go callUrl(url, &wg, ch, chStatus)
	}
	wg.Wait()

	close(chStatus)
	statusMap := make(map[int]int)
	for s := range chStatus {
		q, exists := statusMap[s]
		if exists {
			statusMap[s] = q + 1
		} else {
			statusMap[s] = 1
		}
	}

	printReport(start, requests, statusMap)
}

func callUrl(url string, wg *sync.WaitGroup, ch <-chan interface{}, chStatus chan int) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	chStatus <- response.StatusCode
	<-ch
	wg.Done()
}

func printReport(start time.Time, requests int, statusMap map[int]int) {
	for s, t := range statusMap {
		p := (t * 100) / requests
		fmt.Printf("Status %d ocurred %d times, %d%% of the total\n", s, t, p)
	}

	fmt.Printf("Took %s\n", time.Since(start))
}
