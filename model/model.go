package model

import "gopkg.in/mgo.v2/bson"

type Code struct {
	ID bson.ObjectId `bson:"_id"`
	Title       string
	Description string
	CodeType    string
	CodeString  string
}