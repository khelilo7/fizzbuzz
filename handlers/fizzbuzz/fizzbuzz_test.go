package fizzbuzz_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/maxatome/go-testdeep/helpers/tdhttp"

	"lbc/fizzbuzz/internal/testutils"
)

func TestGetResult(t *testing.T) {
	ta := tdhttp.NewTestAPI(t, testutils.GetRouter())

	ta.Name("test Get result").
		PostJSON("/fizzbuzz",
			json.RawMessage(`{"int1": 2, "int2": 3, "limit": 6, "str1": "fizz", "str2": "buzz"}`)).
		CmpStatus(http.StatusOK).
		CmpJSONBody("1,fizz,buzz,fizz,5,fizzbuzz")
}

// TODO finish test
