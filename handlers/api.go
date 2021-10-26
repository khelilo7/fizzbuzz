package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"lbc/fizzbuzz/handlers/fizzbuzz"
)

type App struct {
	Router *gin.Engine
}

func healthCheck(c *gin.Context) {
	c.Header("Content-Type", "text/plain")
	c.String(http.StatusOK, "Healthy")
}

func (app *App) Init() error {

	app.Router = gin.Default()
	app.Router.GET("/", healthCheck)
	app.Router.POST("/fizzbuzz", fizzbuzz.GetResult)
	return nil
}
