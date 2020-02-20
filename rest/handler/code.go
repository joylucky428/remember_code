package handler

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

	w.Header().Set("Content-Type", "application/json;charset=utf8")
	fmt.Fprintf(w, "{result: success, id: %s}", id)
}

// id 로 코드를 검색하여 반환.
func (h *CodeHandler) GetCode(w http.ResponseWriter, r *http.Request) {
	c, errMsg := h.getCodeByIdFromRequest(r)
	if errMsg != "" {
		fmt.Fprintf(w, errMsg)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&c)
}

func (h *CodeHandler) getCodeByIdFromRequest(r *http.Request) (model.Code, string) {
	// 파라미터 파싱
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return model.Code{}, "{error: No code id}"
	}

	// 코드 가져오기
	idBytes, err := hex.DecodeString(id)
	c, err := h.dbHandler.GetCode(idBytes)
	if err != nil {
		return model.Code{}, "{error occured when fetching code}"
	}

	return c, ""
}

func (h *CodeHandler) DeleteCode(w http.ResponseWriter, r *http.Request) {
	c, errMsg := h.getCodeByIdFromRequest(r)
	if errMsg != "" {
		fmt.Fprintf(w, errMsg)
		return
	}

	err := h.dbHandler.DeleteCode(c)
	if err != nil {
		fmt.Fprintf(w, "{error occured when deleting code %s}", err)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	fmt.Fprintf(w, "{result: 'success'}")
}