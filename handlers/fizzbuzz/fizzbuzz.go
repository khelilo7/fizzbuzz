package fizzbuzz

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"lbc/fizzbuzz/internal/dbmodels"
)

func GetResult(c *gin.Context) {
	log.Info("Hitting Get Result")
	var fizzbuzz dbmodels.Fizzbuzz
	c.ShouldBindJSON(&fizzbuzz)

	var sb strings.Builder

	for i := 1; i <= fizzbuzz.Limit; i++ {
		if i%(fizzbuzz.Int1*fizzbuzz.Int2) == 0 {
			sb.WriteString(fizzbuzz.Str1)
			sb.WriteString(fizzbuzz.Str2)
			sb.WriteString(",")
			continue
		} else if i%int(fizzbuzz.Int1) == 0 {
			sb.WriteString(fizzbuzz.Str1)
			sb.WriteString(",")
			continue
		} else if i%int(fizzbuzz.Int2) == 0 {
			sb.WriteString(fizzbuzz.Str2)
			sb.WriteString(",")
			continue
		}
		sb.WriteString(strconv.FormatInt(int64(i), 10))
		sb.WriteString(",")
	}

	c.JSON(http.StatusOK, sb.String()[:len(sb.String())-1])
}
