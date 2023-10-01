package op

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func (tm *TaskManager) OpenDatabase() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "tasks.db")
    if err != nil {
        return nil, err
    }
    return db, nil
}

func (tm *TaskManager) InsertTaskDB(db *sql.DB, task Task) error {
    insertStatement := `
        INSERT INTO tasks (id, title, description, due_date, completed)
        VALUES (?, ?, ?, ?, ?)
    `
    insertStmt, err := db.Prepare(insertStatement)
    if err != nil {
        return err
    }
    defer insertStmt.Close()
 
    _, err = insertStmt.Exec(task.ID, task.Title, task.Description, task.DueDate, task.Completed)
    if err != nil {
        return err
    }

    return nil
}

func (tm *TaskManager) UpdateTaskDB(db *sql.DB, task Task) error {
    return nil
}

func (tm *TaskManager) DeleteTaskDB(db *sql.DB, id int) error {
    return nil
}

func (tm *TaskManager) LoadTasksFromDatabase(db *sql.DB) error {
    rows, err := db.Query("SELECT id, title, description, due_date, completed FROM tasks")
    if err != nil {
        return err
    }
    defer rows.Close()

    tm.ClearTasks()

    for rows.Next() {
        var task Task
        err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed)
        if err != nil {
            return err
        }
        tm.tasks = append(tm.tasks, task)
        if task.ID > tm.lastID {
            tm.lastID = task.ID
        }
    }

    return nil
}

// func (tm *TaskManager) LoadTasksFromFile(filename string) error {
//     file, err := os.Open(filename)
//     if err != nil {
//         return err
//     }
//     defer file.Close()

//     reader := csv.NewReader(file)
//     lines, err := reader.ReadAll()
//     if err != nil {
//         return err
//     }
//     tm.ClearTasks()

//     for _, line := range lines {
//         if len(line) != 5 {
//             return fmt.Errorf("invalid CSV format")
//         }

//         id, err := strconv.Atoi(line[0])
//         if err != nil {
//             return err
//         }

//         if tm.taskExists(id) {
//             continue
//         }

//         dueDate, err := time.Parse("2006-01-02", line[3])
//         if err != nil {
//             return err
//         }

//         task := Task{
//             ID:          id,
//             Title:       line[1],
//             Description: line[2],
//             DueDate:     dueDate,
//             Completed:   line[4] == "true",
//         }

//         tm.tasks = append(tm.tasks, task)
//         if id > tm.lastID {
//             tm.lastID = id
//         }
//     }

//     return nil
// }

func (tm *TaskManager) taskExists(id int) bool {
    for _, t := range tm.tasks {
        if t.ID == id {
            return true
        }
    }
    return false
}


// func (tm *TaskManager) SaveTasksToFile(filename string) error {
//     file, err := os.Create(filename)
//     if err != nil {
//         return err
//     }
//     defer file.Close()

//     writer := csv.NewWriter(file)
//     defer writer.Flush()

//     for _, task := range tm.tasks {
//         dueDateStr := task.DueDate.Format("2006-01-02")
// 		completedStr := strconv.FormatBool(task.Completed)
//         record := []string{strconv.Itoa(task.ID), task.Title, task.Description, dueDateStr, completedStr}
//         if err := writer.Write(record); err != nil {
//             return err
//         }
//     }

//     return nil
// }
