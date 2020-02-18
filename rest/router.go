package rest

import (
	"github.com/gorilla/mux"
	"remember_code/db"
	"remember_code/rest/handler"
)


func NewRouter(dh db.DatabaseHandler) *mux.Router {
	r := mux.NewRouter()

	codeHandler := &handler.CodeHandler{}
	codeHandler.SetDatabaseHandler(dh)

	codeRouter := r.PathPrefix("/codes").Subrouter()
	codeRouter.Methods("GET").Path("").HandlerFunc(codeHandler.GetCodeList)

	return r
}

