package handle

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/hamanako-palpal/go_smpl_app02/entity"
	"github.com/hamanako-palpal/go_smpl_app02/util"
)

// RootHandler indexページ出力
func RootHandler(w http.ResponseWriter, r *http.Request) {

	var index = template.Must(template.ParseFiles("templates/index.html"))
	err := index.ExecuteTemplate(w, "index.html", "")
	if err != nil {
		log.Fatal("[golang server] internal server error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

// AddUserHandler ユーザ追加
func AddUserHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	user := &entity.User{Name: vars["name"], Pass: vars["pass"]}
	util.DbConn(user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json.Marshal(entity.Request{200, "ok"}))
}
