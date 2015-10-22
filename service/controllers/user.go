package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmanjoura/utm-admin-v0.8/service/models"
	"github.com/mmanjoura/utm-admin-v0.8/service/utilities"
)

type User struct{}

func (u *User) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	db := utilities.GetDB(r)
	user := new(models.User)
	err := user.Get(db, id)
	if err != nil {
		w.WriteHeader(404)
	} else {
		user.Password = ""
		out, _ := json.Marshal(user)
		w.Write(out)
	}
}

func (u *User) Profile(w http.ResponseWriter, r *http.Request) {
	user_id, _ := utilities.GetUserId(r)
	db := utilities.GetDB(r)
	user := new(models.User)
	user.Get(db, user_id)
	user.Password = ""
	out, _ := json.Marshal(user)
	w.Write(out)

}
