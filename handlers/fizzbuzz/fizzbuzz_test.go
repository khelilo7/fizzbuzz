package fizzbuzz_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/td"

	"lbc/fizzbuzz/internal/db"
	"lbc/fizzbuzz/internal/testutils"
)

func TestGetResult(t *testing.T) {
	ta := tdhttp.NewTestAPI(t, testutils.GetRouter())

	ta.Name("test Get result").
		PostJSON("/fizzbuzz",
			json.RawMessage(`{"int1": 2, "int2": 3, "limit": 6, "str1": "fizz", "str2": "buzz"}`)).
		CmpStatus(http.StatusOK).
		CmpJSONBody("1,fizz,buzz,fizz,5,fizzbuzz")

	ta.Name("test Bad input").
		PostJSON("/fizzbuzz",
			json.RawMessage(`{"int1": "2", "int2": 3, "limit": 6, "str1": "fizz", "str2": "buzz"}`)).
		CmpStatus(http.StatusBadRequest).
		CmpJSONBody(td.JSON(`{"status": "Bad input"}`))
}

func TestGetStats(t *testing.T) {
	ta := tdhttp.NewTestAPI(t, testutils.GetRouter())

	ta.Name("test Get Stats").
		Get("/statistics").
		CmpStatus(http.StatusOK)

	// Mock
	// save old one
	var oldGetStats = db.GetStats
	db.GetStats = func(dbc *pg.DB) ([]byte, error) {
		return nil, errors.New("test error")
	}

	ta.Name("test Get Stats failed").
		Get("/statistics").
		CmpStatus(http.StatusInternalServerError).
		CmpJSONBody(td.JSON(`{"status": "Could not get stats"}`))

	db.GetStats = oldGetStats
}
