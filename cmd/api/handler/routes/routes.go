package routes

import (
	"database/sql"

	"github.com/andsanchez/DERES_Back/cmd/api/handler"
	"github.com/andsanchez/DERES_Back/internal/user"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.buildRoutes()
}

func (r *router) buildRoutes() {
	r.r.GET("/ping", handler.PingHandler())
	userRepo := user.NewRepository(r.db)
	userService := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	r.r.GET("/users/login", userHandler.Login())
	r.r.POST("/users/signup", userHandler.SignUp())
}
