package main

import (
	"github.com/andsanchez/DERES_Back/cmd/api/handler/routes"
	"github.com/andsanchez/DERES_Back/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db, err := db.Connect()
	if err != nil {
		panic(err)
	}

	router := routes.NewRouter(r, db)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}
