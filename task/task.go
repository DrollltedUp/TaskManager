package task

import (
	"errors"
	"fmt"
	"time"
)

type Status string

const (
	Todo       Status = "TODO"
	InProgress Status = "In Progress"
	Finished   Status = "Finished"
)

type Task struct {
	id          int
	title       string
	description string
	status      Status
	createdAt   time.Time
	updatedAt   time.Time
}

type Tasks []Task

func (tasks *Tasks) createTask(title string, description string) error {
	task := Task{
		id:          0,
		title:       title,
		description: description,
		status:      "TODO",
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}

	*tasks = append(*tasks, task)

	return nil
}

func (tasks *Tasks) IsValidationTask(index int) error {
	if index == 0 || index > len(*tasks) {
		err := errors.New("Now Validate for Tasks")
		fmt.Println(err)
		return err
	}

	return nil
}

func (tasks *Tasks) DeleteTask(index int) error {
	t := *tasks
	if error := t.IsValidationTask(index); error != nil {
		fmt.Println(error)
		return error
	}
	*tasks = append(t[:index], t[:index-1]...)
	return nil
}

func (tasks *Tasks) UpdateStatusTask(index int) error {
	t := *tasks
	if error := t.IsValidationTask(index); error != nil {
		fmt.Println(error)
		return error
	}
}
