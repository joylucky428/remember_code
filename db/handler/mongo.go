package handler

import (
	"gopkg.in/mgo.v2"
	"log"
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
	log.Println("mongo : GetCodeList")

	s := h.getFreshSession()
	defer s.Close()
	var codes []model.Code
	err := s.DB(DATABASE_NAME).C(CODE_COLL_NAME).Find(nil).All(&codes)
	return codes, err
}