package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"remember_code/db"
	"remember_code/rest/handler"
)


func NewRouter(dh db.DatabaseHandler) http.Handler {
	r := mux.NewRouter()

	codeHandler := &handler.CodeHandler{}
	codeHandler.SetDatabaseHandler(dh)

	codeRouter := r.PathPrefix("/codes").Subrouter()
	codeRouter.Methods("GET").Path("").HandlerFunc(codeHandler.GetCodeList)
	codeRouter.Methods("POST").Path("").HandlerFunc(codeHandler.AddCode)
	codeRouter.Methods("GET").Path("/{id}").HandlerFunc(codeHandler.GetCode)

	return corsMiddleware(r)
}


func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token, Authorization")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
