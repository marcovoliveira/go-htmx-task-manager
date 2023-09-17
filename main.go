package main

import (
	"bufio"
	"fmt"
	utils "go-task-manager-cli/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	tm := NewTaskManager()

	err := tm.LoadTasksFromFile("tasks.csv")
    if err != nil {
        fmt.Println("Error loading tasks:", err)
    }


	reader := bufio.NewReader(os.Stdin)


	for {
		fmt.Println("Task Manager")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter task title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			fmt.Print("Enter task description: ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)
			fmt.Print("Enter due date (YYYY-MM-DD): ")
			dateStr, _ := reader.ReadString('\n')
			dateStr = strings.TrimSpace(dateStr)

			dueDate, err := utils.ParseDate(dateStr)
			if err != nil {
				fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
				continue
			}
			
			tm.AddTask(title, description, dueDate)
		case "2":
			tm.ListTasks()
		case "3":
			fmt.Print("Enter task ID to mark as completed: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)

			taskID, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid task ID. Please enter a valid number.")
				continue
			}

			err = tm.MarkTaskCompleted(taskID)
			if err != nil {
				fmt.Println("Task not found.")
			} else {
				fmt.Println("Task marked as completed.")
			}
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}