package models

import (
    "database/sql"
    //"github.com/labstack/echo"
)

// Task is a struct containing Task data
type Task struct {
    ID int `json:"id"`
    NAME string `json:"name"`
}

// TaskCollection is collection of Tasks
type TaskCollections struct {
    Tasks []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollections {
    sql := "SELECT * FROM tasks"

    rows, err := db.Query(sql)

    if err != nil {
        panic(err)
    }

    defer rows.Close()

    result := TaskCollections{}

    for rows.Next() {
        task := Task{}
        err2 := rows.Scan(&task.ID, &task.NAME)

        if err2 != nil {
            panic(err2)
        }

        result.Tasks = append(result.Tasks, task)
    }

    return result
}

func PutTask(db *sql.DB, name string) (int64, error) {
    sql := "INSERT INTO tasks(name) VALUES(?)"

    stmt, err := db.Prepare(sql)

    if err != nil {
        panic(err)
    }

    defer stmt.Close()

    result, err2 := stmt.Exec(name)

    if err2 != nil {
        panic(err2)
    }

    return result.LastInsertId()
}
