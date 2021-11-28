package db_test

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"

	"lbc/fizzbuzz/internal/db"
)

func TestGetStats(t *testing.T) {
	data, err := db.GetStats(db.NewDBConn())
	td.CmpNotNil(t, data)
	td.CmpNoError(t, err)
}
