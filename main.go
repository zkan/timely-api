package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zkan/timely-api/internal/middleware"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("port", "1323")
	viper.SetDefault("cors.allow_origin", "*")
}

func main() {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(middleware.Headers(viper.GetString("cors.allow_origin")))

	r.HandleFunc("/healths", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + viper.GetString("port"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
