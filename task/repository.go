package task

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"encoding/json"
	"github.com/gorilla/mux"
)

type Task struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Username  string    `json:"username"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
}

type TaskCreationRequest struct {
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Username  string    `json:"username"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
}

type Tasks []Task

func HandleRequest(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			stmt := "SELECT id, name, category, username, started_at, ended_at FROM tasks"
			rows, err := db.Query(stmt)
			if err != nil {
				log.Print(err)
			}
			defer rows.Close()
			tasks := make(Tasks, 0)

			for rows.Next() {
				var task Task
				if err := rows.Scan(&task.ID, &task.Name, &task.Category, &task.Username, &task.StartedAt, &task.EndedAt); err != nil {
					log.Fatal(err)
				}
				tasks = append(tasks, task)
			}

			js, err := json.Marshal(tasks)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		case "POST":
			var t TaskCreationRequest

			if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			stmt := `INSERT into tasks (name, category, username, started_at, ended_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

			var lastInsertId int64
			err := db.QueryRow(stmt, t.Name, t.Category, t.Username, t.StartedAt, t.EndedAt).Scan(&lastInsertId)
			if err != nil {
				panic(err)
			}

			task := Task{
				ID:        lastInsertId,
				Name:      t.Name,
				Category:  t.Category,
				Username:  t.Username,
				StartedAt: t.StartedAt,
				EndedAt:   t.EndedAt,
			}

			js, err := json.Marshal(task)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		}
	}
}

func Delete(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err == nil {
			log.Print(err)
		}

		stmt := `DELETE FROM tasks WHERE id = $1`
		db.Exec(stmt, id)

		json.NewEncoder(w).Encode(http.StatusOK)
	}
}
