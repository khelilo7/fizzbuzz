package db_test

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"

	"lbc/fizzbuzz/internal/db"
)

func TestAggReqCount(t *testing.T) {
	req, err := db.AggReqCount(db.NewDBConn())
	td.CmpNotNil(t, req)
	td.CmpNoError(t, err)
}

func TestAggFieldCount(t *testing.T) {
	// error
	req, err := db.AggFieldCount(db.NewDBConn(), "")
	td.CmpNil(t, req)
	td.CmpError(t, err)

	// no error
	req, err = db.AggFieldCount(db.NewDBConn(), "int1")
	td.CmpNotNil(t, req)
	td.CmpNoError(t, err)
}
