package testutils

import (
	"fmt"

	"github.com/gin-gonic/gin"

	api "lbc/fizzbuzz/handlers"
)

func GetRouter() *gin.Engine {
	app := new(api.App)
	err := app.Init()
	if err != nil {
		fmt.Println(err)
	}
	return app.Router
}
