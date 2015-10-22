package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goincremental/negroni-sessions"
	"github.com/mmanjoura/utm-admin-v0.8/service/models"
	"github.com/mmanjoura/utm-admin-v0.8/service/utilities"
)

type Credentials struct {
	Email    string
	Password string
}

type Auth struct{}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	credentials := new(Credentials)
	err := decoder.Decode(&credentials)
	if err != nil {
		panic(err)
	}

	db := utilities.GetDB(r)
	user := new(models.User)
	err = user.Authenticate(db, credentials.Email, credentials.Password)
	if err == nil {
		session := sessions.GetSession(r)
		session.Set("user_id", user.ID.Hex())
		session.Set("user_company", user.Company)
		session.Set("user_email", user.Email)
		w.WriteHeader(202)

	} else {
		w.WriteHeader(404)
	}
}

func (a *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	session := sessions.GetSession(r)
	user_id := session.Get("user_id")
	fmt.Println(user_id)
	if user_id == nil {
		w.WriteHeader(403)
		http.Redirect(w, r, "/", 403)

	} else {
		session.Delete("user_id")
		http.Redirect(w, r, "/", 202)
	}

}

func (a *Auth) User(w http.ResponseWriter, r *http.Request) {
	db := utilities.GetDB(r)
	session := sessions.GetSession(r)
	user_id := session.Get("user_id")
	fmt.Println(user_id)
	if user_id == nil {
		w.WriteHeader(403)

	} else {
		user := new(models.User)
		user.Get(db, user_id.(string))
		fmt.Println(user)
		outData, _ := json.Marshal(user)
		w.Write(outData)
	}

}

func (a *Auth) Users(w http.ResponseWriter, r *http.Request) {
	db := utilities.GetDB(r)
	session := sessions.GetSession(r)
	user_company := session.Get("user_company")
	fmt.Println(user_company)
	if user_company == nil {
		w.WriteHeader(403)

	} else {
		user := new(models.User)
		users := user.GetCompanyUsers(db, user_company.(string))
		fmt.Println(users)
		outData, _ := json.Marshal(users)
		w.Write(outData)
	}

}

func (a *Auth) Uuids(w http.ResponseWriter, r *http.Request) {
	db := utilities.GetDB(r)
	session := sessions.GetSession(r)
	user_company := session.Get("user_company")
	fmt.Println(user_company)
	if user_company == nil {
		w.WriteHeader(403)

	} else {
		user := new(models.User)
		uuids := user.GetUserUuids(db, user_company.(string))
		fmt.Println(uuids)
		outData, _ := json.Marshal(uuids)
		w.Write(outData)
	}

}

func (a *Auth) UserUEs(w http.ResponseWriter, r *http.Request) {
	db := utilities.GetDB(r)
	session := sessions.GetSession(r)
	user_email := session.Get("user_email")
	fmt.Println(user_email)
	if user_email == nil {
		w.WriteHeader(403)

	} else {
		user := new(models.User)
		ues := user.GetUEs(db, user_email.(string))
		fmt.Println(ues)
		outData, _ := json.Marshal(ues)
		w.Write(outData)

		// fmt.Printf("%s Dave Chany magic file dump", spew.Sdump(outData))
	}

}

func (a *Auth) Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	data := map[string]string{"company": "", "firstName": "", "last Name": "", "userName": "", "email": "", "password": ""}
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	db := utilities.GetDB(r)
	user := new(models.User)

	//temporary
	// uuid := new(models.Uuid)
	// uuid.Insert(db)
	user.NewUser(db, data["company"], data["firstName"], data["lastName"], data["userName"], data["email"], data["password"])
	fmt.Println(user)

}
