package handle

import (
	"log"
	"net/http"
	"text/template"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {

	var index = template.Must(template.ParseFiles("templates/index.html"))
	err := index.ExecuteTemplate(w, "index.html", "")
	if err != nil {
		log.Fatal("[golang server] internal server error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {

}
