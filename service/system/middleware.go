package system

import (
	"github.com/mmanjoura/utm-admin-v0.8/service/utilities"
	"gopkg.in/mgo.v2"
	"net/http"
)

var logTag string = "UTM"

func MgoMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	session, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		panic(err)
	}

	reqSession := session.Clone()
	defer reqSession.Close()
	dbs := reqSession.DB("utm-db")
	utilities.SetDB(r, dbs)
	next(rw, r)
}
