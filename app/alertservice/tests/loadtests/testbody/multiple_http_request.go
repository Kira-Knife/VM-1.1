package testbody

import (
	"loadtests/entity"
	"loadtests/http_request"
	"sync"
	"time"
)

func multiple_http_request(countFetch int, method string, url string, jsonData []byte) (time.Duration, []entity.RequestResult, error) {
	var wg sync.WaitGroup
	results := make(chan entity.RequestResult, countFetch)

	startTime := time.Now()

	for i := 0; i < countFetch; i++ {
		wg.Add(1)
		go http_request.ExecHttpRequestWithTime(&wg, method, url, jsonData, results)

	}

	wg.Wait()
	close(results)

	endTime := time.Now()
	totalTime := endTime.Sub(startTime)
	var reqRes = make([]entity.RequestResult, 0, countFetch)
	for r := range results {
		reqRes = append(reqRes, r)
	}

	return totalTime, reqRes, nil

}
