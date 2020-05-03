package main

import (
	"fmt"
	"net/http"

	"github.com/zkan/timely-api/internal/middleware"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)


func main() {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(middleware.Headers(viper.GetString("cors.allow_origin")))

	r.HandleFunc("/healths", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
}
