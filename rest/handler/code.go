package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"remember_code/db"
)

type CodeHandler struct {
	dbHandler db.DatabaseHandler
}

func (h *CodeHandler) GetCodeList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getCodeList")
	cl, err := h.dbHandler.GetCodeList()
	if err != nil {
		log.Fatal("error when get code list from db.")
	}

	fmt.Println(cl)
	json.NewEncoder(w).Encode(&cl)
}

func (h *CodeHandler) SetDatabaseHandler(dh db.DatabaseHandler) {
	h.dbHandler = dh
}