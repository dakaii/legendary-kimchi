package main

import (
	"fmt"
	"graphyy/controller"
	"graphyy/database"
	"graphyy/internal/envvar"
	"graphyy/repository"
	"net/http"
)

func main() {
	context, db := database.InitDatabase()
	repos := repository.InitRepositories(context, db)
	controllers := controller.InitControllers(repos)
	schema := controller.Schema(controllers)

	http.Handle("/graphql", controller.GraphqlHandlfunc(schema))

	fmt.Println("server is started at: http://localhost:/" + envvar.ServerPort() + "/")
	fmt.Println("graphql api server is started at: http://localhost:" + envvar.ServerPort() + "/graphql")
	http.ListenAndServe(":"+envvar.ServerPort(), nil)
}
