package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func (tm *TaskManager) LoadTasksFromFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    lines, err := reader.ReadAll()
    if err != nil {
        return err
    }

    for _, line := range lines {
        if len(line) != 5 {
            return fmt.Errorf("invalid CSV format")
        }

        id, err := strconv.Atoi(line[0])
        if err != nil {
            return err
        }

        dueDate, err := time.Parse("2006-01-02", line[3])
        if err != nil {
            return err
        }


        task := Task{
            ID:          id,
            Title:       line[1],
            Description: line[2],
            DueDate:     dueDate,
            Completed:   line[4] == "true",
        }
        tm.tasks = append(tm.tasks, task)
        if id > tm.lastID {
            tm.lastID = id
        }
    }

    return nil
}

func (tm *TaskManager) SaveTasksToFile(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, task := range tm.tasks {
        dueDateStr := task.DueDate.Format("2006-01-02")
		completedStr := strconv.FormatBool(task.Completed)
        record := []string{strconv.Itoa(task.ID), task.Title, task.Description, dueDateStr, completedStr}
        if err := writer.Write(record); err != nil {
            return err
        }
    }

    return nil
}
