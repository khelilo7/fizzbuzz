package fizzbuzz

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"lbc/fizzbuzz/internal/db"
	"lbc/fizzbuzz/internal/dbmodels"
	"lbc/fizzbuzz/internal/utils"
)

func GetResult(c *gin.Context) {
	log.Info("Hitting Get Result")
	var fizzbuzz dbmodels.Fizzbuzz
	c.ShouldBindJSON(&fizzbuzz)

	dbc := db.NewDBConn()
	_, err := dbc.Model(&fizzbuzz).Insert()
	if err != nil {
		log.Errorf("Could not create new annonce, err %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Could not create new annonce"})
		return
	}

	c.JSON(http.StatusOK, utils.BuildFizzBuzzString(fizzbuzz))
}

func GetStats(c *gin.Context) {
	log.Info("Hitting Get Stats")
	dbc := db.NewDBConn()

	jsonData, err := db.GetStats(dbc)
	if err != nil {
		log.Error("Could not get stats, err %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Could not get stats"})
		return
	}

	c.Data(http.StatusOK, "application/json", jsonData)
}
