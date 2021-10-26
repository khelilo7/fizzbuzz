package utils_test

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"

	"lbc/fizzbuzz/internal/utils"
)

func TestGetEnv(t *testing.T) {
	td.Cmp(t, utils.Getenv("something", "test"), "test")
}
