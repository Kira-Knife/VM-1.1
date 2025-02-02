package testbody

import (
	"loadtests/statistics"
	"log"
	"net/http"
)

func TestInsertAlerts(host string, countFetch int) {
	method := http.MethodPost
	url := host + "/api/v1/alerts"

	var jsonDataInsert = []byte(`{
    "receiver": "test",
    "status": "firing",
    "alerts": [
      {
        "status": "firing",
        "labels": {
          "alertname": "TestAlert_1001",
          "instance": "Grafana"
        },
        "annotations": {
          "summary": "Notification test"
        },
        "startsAt": "2024-12-03T07:14:03.470708272Z",
        "endsAt": "0001-01-01T00:00:00Z",
        "generatorURL": "",
        "fingerprint": "57c6d9296de2ad39",
        "silenceURL": "http://localhost:3000/alerting/silence/new?alertmanager=grafana&matcher=alertname%3DTestAlert&matcher=instance%3DGrafana",
        "dashboardURL": "",
        "panelURL": "",
        "values": null,
        "valueString": "[ metric='foo' labels={instance=bar} value=10 ]"
      }
    ],
    "groupLabels": {
      "alertname": "TestAlert",
      "instance": "Grafana"
    },
    "commonLabels": {
      "alertname": "TestAlert",
      "instance": "Grafana"
    },
    "commonAnnotations": {
      "summary": "Notification test"
    },
    "externalURL": "http://localhost:3000/",
    "version": "1",
    "groupKey": "test-57c6d9296de2ad39-1733210043",
    "truncatedAlerts": 0,
    "orgId": 1,
    "title": "[FIRING:1] TestAlert Grafana ",
    "state": "alerting",
    "message": "**Firing**\n\nValue: [no value]\nLabels:\n - alertname = TestAlert\n - instance = Grafana\nAnnotations:\n - summary = Notification test\nSilence: http://localhost:3000/alerting/silence/new?alertmanager=grafana&matcher=alertname%3DTestAlert&matcher=instance%3DGrafana\n"
}`)

	totalTime, reqRes, err := multiple_http_request(countFetch, method, url, jsonDataInsert)
	if err != nil {
		log.Fatal(err)
		return
	}

	statistics.PrintStatistics("TestAlertInsert", totalTime, reqRes)
}
