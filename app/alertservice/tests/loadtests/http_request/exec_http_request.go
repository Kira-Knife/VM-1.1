package http_request

import (
	"bytes"
	"crypto/tls"
	"loadtests/entity"
	"net/http"
	"sync"
	"time"
)

func ExecHttpRequestWithTime(wg *sync.WaitGroup, method string, url string, jsonData []byte, results chan<- entity.RequestResult) {
	defer wg.Done()

	start := time.Now()

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		results <- entity.RequestResult{
			Error: err,
		}
		return
	}

	if jsonData != nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("accept", "application/json")
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transport}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	duration := time.Since(start)
	r := entity.RequestResult{
		ExecutionTime: duration,
		StatusCode:    resp.StatusCode,
	}

	results <- r
}
