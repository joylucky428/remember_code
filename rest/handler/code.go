package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"remember_code/db"
	"remember_code/model"
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

// 코드 추가
func (h *CodeHandler) AddCode(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddCode api..")

	c := model.Code{}

	// request json data 파싱
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: error occured while parsing code json data %s}", err)
		return
	}

	// Code 저장
	id, err := h.dbHandler.AddCode(c)
	if nil != err {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: error occured while persisting code %d %s}", id, err)
		return
	}

	fmt.Fprintf(w, "{result: success, id: %s}", id)
}