package main

import "loadtests/testbody"

func main() {
	numberOfRequests := 100                           // число запросов
	host := "https://localhost:8787"                  // хост на который пойдут запросы
	testbody.TestInsertAlerts(host, numberOfRequests) // запросы на вставку данных
	testbody.TestSelectAlerts(host, numberOfRequests) // запросы на получение данных
}
