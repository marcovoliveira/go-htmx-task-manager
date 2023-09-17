package main

import (
	"fmt"
	"time"
)

type TaskManager struct {
	tasks []Task
	lastID int
}

type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     time.Time
	Completed   bool
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (tm *TaskManager) 	AddTask(title, description string, dueDate time.Time) {
	tm.lastID++
	task := Task{
		ID:          tm.lastID,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Completed:   false,
	}
	tm.tasks = append(tm.tasks, task)
	err := tm.SaveTasksToFile("tasks.csv")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

func (tm *TaskManager) ListTasks() {
	if len(tm.tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("Tasks:")
	for _, task := range tm.tasks {
		status := "Pending"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("ID: %d\nTitle: %s\nDescription: %s\nDue Date: %s\nStatus: %s\n\n",
			task.ID, task.Title, task.Description, task.DueDate.Format("2006-01-02"), status)
	}
}

func (tm *TaskManager) MarkTaskCompleted(id int) error {
	for i, task := range tm.tasks {
		if task.ID == id {
			tm.tasks[i].Completed = true
			err := tm.SaveTasksToFile("tasks.csv")
            if err != nil {
                return err
            }
			return nil
		}
	}
	
	return fmt.Errorf("task not found")
}