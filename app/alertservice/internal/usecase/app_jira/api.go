package appjira

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type TaskJiraRequest struct {
	TaskTitle  string `json:"task_title" example:"Task Title"`
	Status     string `json:"task_status" example:"Открыт"`
	Assigned   string `json:"assigned" example:"Иван Иванов"`
	Owner      string `json:"owner" example:"Петр Петров"`
	IncidentId int64  `json:"incident_id" example:"1"`
}

type TaskStatus struct {
	Status string `json:"status"`
}

// CreateTaskJira - обращается к API для создания задачи Jira
func (a *AppJiraAPI) CreateTaskJira(ctx context.Context, taskReq TaskJiraRequest) (uuid.UUID, error) {
	url := a.url + "/api/v1/task"
	jsonData, err := json.Marshal(taskReq)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return uuid.UUID{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var taskUUID uuid.UUID
	err = json.NewDecoder(resp.Body).Decode(&taskUUID)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return taskUUID, nil
}

// UpdateStatusTaskJira - обращается к API для обновления статуса задачи Jira
func (a *AppJiraAPI) UpdateStatusTaskJira(ctx context.Context, incidentID int64, newStatus string) error {
	url := fmt.Sprintf("%s/api/v1/task/incident/%d", a.url, incidentID)
	status := TaskStatus{Status: newStatus}
	jsonData, err := json.Marshal(status)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
