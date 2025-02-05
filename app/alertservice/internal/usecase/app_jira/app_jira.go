package appjira

import (
	"net/http"
	"time"
)

type AppJiraAPI struct {
	client *http.Client
	url    string // http://localhost:8786
}

func New(urlAppJira string) *AppJiraAPI {
	return &AppJiraAPI{
		url: urlAppJira,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}
