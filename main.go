package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hamanako-palpal/go_smpl_app02/handle"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handle.RootHandler)
	r.HandleFunc("/add/{name}/{pass}", handle.AddUserHandler)
	s := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	s.ListenAndServe()
}
