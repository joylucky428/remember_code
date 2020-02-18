package db

import (
	"remember_code/model"
)

type DatabaseHandler interface {
	GetCodeList() ([]model.Code, error)
	AddCode(model.Code) ([]byte, error)
	GetCode(id []byte) (model.Code, error)
}