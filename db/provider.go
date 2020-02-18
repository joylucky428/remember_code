package db

import "remember_code/db/handler"

func NewDatabaseHandler(dbType, connection string) (DatabaseHandler, error){
	switch dbType {
	case "mongo":
		return handler.NewMongoHandler(connection)
	}
	return nil, nil
}