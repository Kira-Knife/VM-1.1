package entity

import "github.com/google/uuid"

type TaskJira struct {
	TaskId     int64     `json:"task_id" example:"1"`
	JiraTaskId uuid.UUID `json:"jira_task_id"`
	TaskTitle  string    `json:"task_title" example:"Task Title"`
	Status     string    `json:"task_status" example:"Открыт"`
	Assigned   string    `json:"assigned" example:"Иван Иванов"`
	Owner      string    `json:"owner" example:"Петр Петров"`
	IncidentId int64     `json:"incident_id" example:"1"`
}
