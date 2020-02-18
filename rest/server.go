package rest

import (
	"log"
	"net/http"
	"remember_code/db"
)

func ServeAPI(endpoint string, dh db.DatabaseHandler) error {
	log.Println("serve api..")
	log.Println("endpoint : ", endpoint)

	r := NewRouter(dh)
	err := http.ListenAndServe(endpoint, r)

	return err
}