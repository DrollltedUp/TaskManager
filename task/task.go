package task

import "time"

type CRUD interface {
	AddTask()
}

type Status string

const (
	Todo       Status = "TODO"
	InProgress Status = "In Progress"
	Finished   Status = "Finished"
)

type Priority string

const (
	Low    Priority = "Low"
	Medium Priority = "Medium"
	High   Priority = "High"
)

type Task struct {
	id          int
	title       string
	description string
	status      Status
	priority    Priority
	createdAt   time.Time
	updatedAt   time.Time
}

fmt.print