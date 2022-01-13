package main

import (
	"final-project/db"
	"final-project/docs"
	"final-project/router"
)

const PORT = ":8080"

func init() {
	db.StartDB("localhost", 5432, "postgres", "postgres", "todo")
}

func main() {
	docs.SwaggerInfo.Title = "Swagger To Do Rest API"
	docs.SwaggerInfo.Description = "Documentation of Todo Rest API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost" + PORT
	docs.SwaggerInfo.Schemes = []string{"http"}
	r := router.StartApp()
	r.Run(PORT)
	db.GetDB().Close()
}
