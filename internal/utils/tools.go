package utils

import (
	"os"
	"strconv"
	"strings"

	"lbc/fizzbuzz/internal/dbmodels"
)

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func BuildFizzBuzzString(fizzbuzz dbmodels.Fizzbuzz) string {
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
	return sb.String()[:len(sb.String())-1]
}
