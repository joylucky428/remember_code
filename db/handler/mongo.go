package handler

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"remember_code/model"
)

const (
	DATABASE_NAME = "remember_code"
	CODE_COLL_NAME = "codes"
)

type MongoHandler struct {
	session *mgo.Session
}

func NewMongoHandler(connection string) (*MongoHandler, error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}

	handler := &MongoHandler{
		session: session,
	}
	return handler, err
}

func (h *MongoHandler) getFreshSession() *mgo.Session {
	return h.session.Copy()
}


func (h *MongoHandler) GetCodeList() ([]model.Code, error) {
	fmt.Println("mongo...?")
	var cl []model.Code

	newCode := model.Code{
		Title:       "Sample Code",
		Description: "Golang code",
		CodeType:    "Go",
		CodeString:  "fmt.Println(\"hello world\")",
	}

	cl = append(cl, newCode)

	return cl, nil
}