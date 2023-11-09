package main

import (
	"walmartTest/pkg/repository"
	"walmartTest/pkg/router"
)

func main() {
	//Init MongoDatabase
	repository.InitMongoDatabase()

	//Init Routers
	router.InitRouter()
}
