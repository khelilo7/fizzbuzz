package db_test

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"

	"lbc/fizzbuzz/internal/db"
)

func TestNewDBConn(t *testing.T) {
	td.CmpNotNil(t, db.NewDBConn())
}
