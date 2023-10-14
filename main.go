package main

import (
	"fmt"
	op "go-task-manager-cli/src"
	util "go-task-manager-cli/utils"
	"html/template"
	"log"
	"net/http"
	"time"
)

const Port = ":8000"

func main() {
	fmt.Printf("Server running on port http://localhost%s...", Port)
	tm := op.NewTaskManager()
	//err := tm.LoadTasksFromFile("tasks.csv")
	db, err := tm.OpenDatabase()
	tm.LoadTasksFromDatabase(db)

	if err != nil {
		fmt.Println("Error loading tasks:", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		displayTasks(w, tm)
	})

	http.HandleFunc("/task/", func(w http.ResponseWriter, r *http.Request) {
		taskHandler(w, r, tm)
		tmpl := template.Must(template.ParseFiles("layouts/index.html"))
		tasks := map[string][]op.Task{
			"Tasks": tm.GetTasks(),
		}
		tmpl.Execute(w, tasks)
	})

	log.Fatal(http.ListenAndServe(Port, nil))
}

func displayTasks(w http.ResponseWriter, tm *op.TaskManager) {
	tmpl := template.Must(template.ParseFiles("layouts/index.html"))
	tasks := map[string][]op.Task{
		"Tasks": tm.GetTasks(),
	}
	tmpl.Execute(w, tasks)
}

func taskHandler(w http.ResponseWriter, r *http.Request, tm *op.TaskManager) {
	switch r.Method {
	case http.MethodPost:
		addTask(w, r, tm)
	case http.MethodDelete:
		deleteTask(w, r, tm)
	default:
		http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
	}
}

func addTask(w http.ResponseWriter, r *http.Request, tm *op.TaskManager) {
	title := r.PostFormValue("title")
	description := r.PostFormValue("description")
	dueDate, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	tm.AddTask(title, description, dueDate)
}

func deleteTask(w http.ResponseWriter, r *http.Request, tm *op.TaskManager) {
	urlPath := r.URL.Path
	id := urlPath[len("/task/"):]
	tm.DeleteTask(util.ParseInt(id))
}
