package utils_test

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"

	"lbc/fizzbuzz/internal/dbmodels"
	"lbc/fizzbuzz/internal/utils"
)

var fizzbuzz = dbmodels.Fizzbuzz{
	Int1:  2,
	Int2:  3,
	Limit: 6,
	Str1:  "fizz",
	Str2:  "buzz",
}

func TestGetEnv(t *testing.T) {
	td.Cmp(t, utils.Getenv("something", "test"), "test")
}

func TestBuildFizzBuzzString(t *testing.T) {
	td.Cmp(t, utils.BuildFizzBuzzString(fizzbuzz), "1,fizz,buzz,fizz,5,fizzbuzz")
}
