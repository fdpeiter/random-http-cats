package main

import (
	"github.com/fdpeiter/random-http-cats/app/routes"
)

func main() {
	r := routes.SetupRouter()
	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
