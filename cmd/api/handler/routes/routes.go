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
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildRoutes() {
	r.r.GET("/ping", handler.PingHandler())
	userRepo := user.NewRepository(r.db)
	userService := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	r.rg.GET("/users/login", userHandler.Login())
	r.rg.POST("/users/signup", userHandler.SignUp())
}
