package main

import (
	"fmt"
	op "go-task-manager-cli/src"
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

	generateIndex := func(w http.ResponseWriter, r *http.Request) {
		//err := tm.LoadTasksFromFile("tasks.csv")
		db, err := tm.OpenDatabase()
		tm.LoadTasksFromDatabase(db)
		
		if err != nil {
			fmt.Println("Error loading tasks:", err)
		}
		tmpl := template.Must(template.ParseFiles("layouts/index.html"))
		tasks := map[string][]op.Task{
			"Tasks": tm.GetTasks(),
		}		
		tmpl.Execute(w, tasks)
	}

	saveTask := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		description := r.PostFormValue("description")
		tmpl := template.Must(template.ParseFiles("layouts/index.html"))
		dueDate, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
		tm.AddTask(title, description, dueDate)
		tasks := map[string][]op.Task{
			"Tasks": tm.GetTasks(),
		}		
		tmpl.Execute(w, tasks)
		//tmpl.ExecuteTemplate(w, "task-list-element", Task{Title: title, Description: description})
	}

	http.HandleFunc("/", generateIndex)
	http.HandleFunc("/add-task/", saveTask)
	log.Fatal(http.ListenAndServe(Port, nil))
}