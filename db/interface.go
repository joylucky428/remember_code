package db

import (
	"remember_code/model"
)

type DatabaseHandler interface {
	GetCodeList() ([]model.Code, error)
}

