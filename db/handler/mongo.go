package handler

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

// 새로운 코드를 추가하고 ID를 반환한다.
func (h *MongoHandler) AddCode(c model.Code) ([]byte, error) {
	fmt.Println("mongo add....")

	s := h.getFreshSession()
	defer s.Close()
	if !c.ID.Valid() {
		c.ID = bson.NewObjectId()
	}
	return []byte(c.ID), s.DB(DATABASE_NAME).C(CODE_COLL_NAME).Insert(c)
}

// ID로 코드 검색
func (h *MongoHandler) GetCode(id []byte) (model.Code, error) {
	s := h.getFreshSession()
	defer s.Close()
	c := model.Code{}
	err := s.DB(DATABASE_NAME).C(CODE_COLL_NAME).FindId(bson.ObjectId(id)).One(&c)
	return c, err
}

// ID로 코드 삭제
func (h *MongoHandler) DeleteCode(c model.Code) error {
	s := h.getFreshSession()
	defer s.Close()
	err := s.DB(DATABASE_NAME).C(CODE_COLL_NAME).Remove(c)
	return err
}