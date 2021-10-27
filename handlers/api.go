package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"

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
	app.Router.Use(stats.RequestStats())

	app.Router.GET("/", healthCheck)
	app.Router.POST("/fizzbuzz", fizzbuzz.GetResult)
	app.Router.GET("/statistics", fizzbuzz.GetStats)
	app.Router.GET("/api_stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})
	return nil
}
