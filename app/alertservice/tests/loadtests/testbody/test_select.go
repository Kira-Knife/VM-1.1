package testbody

import (
	"loadtests/statistics"
	"log"
	"net/http"
)

func TestSelectAlerts(host string, countFetch int) {
	method := http.MethodGet
	url := host + "/api/v1/alerts"

	totalTime, reqRes, err := multiple_http_request(countFetch, method, url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	statistics.PrintStatistics("TestSelectAlerts", totalTime, reqRes)
}
