package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/zkan/timely-api/internal/middleware"
	"github.com/zkan/timely-api/task"
)

func init() {
	viper.SetDefault("port", "1323")
	viper.SetDefault("cors.allow.origin", "*")
	viper.SetDefault("db.conn.string", "host=localhost user=postgres password=mysecretpassword dbname=postgres sslmode=disable")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func newDBClient(connStr string) (*sql.DB, error) {
	return sql.Open("postgres", connStr)
}

func main() {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(middleware.Headers(viper.GetString("cors.allow.origin")))

	r.HandleFunc("/healths", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	db, err := newDBClient(viper.GetString("db.conn.string"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r.HandleFunc("/tasks", task.HandleRequest(db)).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/tasks/{id}", task.Delete(db)).Methods(http.MethodDelete)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + viper.GetString("port"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
