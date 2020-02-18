package main

import (
	"fmt"
	"log"
	"remember_code/config"
	"remember_code/db"
	"remember_code/rest"
	"strconv"
)

func main()  {
	fmt.Println("start..")

	endPoint := ":" + strconv.Itoa(config.ApiServerPort)

	dbType := config.DbType
	dbConnection := config.DbConnection
	dh, err := db.NewDatabaseHandler(dbType, dbConnection)
	if err != nil {
		log.Fatal("error when get new database handler", err)
	}

	if err := rest.ServeAPI(endPoint, dh); err != nil {
		log.Fatal("error when serve api")
	}
}
