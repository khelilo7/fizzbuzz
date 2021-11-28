package db

import (
	"encoding/json"

	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"

	"lbc/fizzbuzz/internal/dbmodels"
)

var GetStats = func(dbc *pg.DB) ([]byte, error) {
	req, err := AggReqCount(dbc)
	if err != nil {
		log.Errorf("Could not query requests, err %v", err)
		return nil, err
	}
	int1, err := AggFieldCount(dbc, "int1")
	if err != nil {
		log.Errorf("Could not query requests, err %v", err)
		return nil, err
	}

	int2, err := AggFieldCount(dbc, "int2")
	if err != nil {
		log.Errorf("Could not query requests, err %v", err)
		return nil, err
	}

	int12, err := AggFieldCount(dbc, "int1", "int2")
	if err != nil {
		log.Errorf("Could not query requests, err %v", err)
		return nil, err
	}

	limit, err := AggFieldCount(dbc, "limite")
	if err != nil {
		log.Errorf("Could not query requests, err %v", err)
		return nil, err
	}

	str1, err := AggFieldCount(dbc, "str1")
	if err != nil {
		log.Errorf("Could not query requests, err %v", err)
		return nil, err
	}

	str2, err := AggFieldCount(dbc, "str2")
	if err != nil {
		log.Errorf("Could not query requests, err %v", err)
		return nil, err
	}

	str12, err := AggFieldCount(dbc, "str1", "str2")
	if err != nil {
		log.Errorf("Could not query requests, err %v", err)
		return nil, err
	}

	strInt1, err := AggFieldCount(dbc, "str1", "int1")
	if err != nil {
		log.Errorf("Could not query requests, err %v", err)
		return nil, err
	}

	strInt2, err := AggFieldCount(dbc, "str2", "int2")
	if err != nil {
		log.Errorf("Could not query requests, err %v", err)
		return nil, err
	}

	data := struct {
		Req     dbmodels.Request `json:"Most used request"`
		Int1    dbmodels.Request `json:"Most used int1"`
		Int2    dbmodels.Request `json:"Most used int2"`
		Limit   dbmodels.Request `json:"Most used limit"`
		Str1    dbmodels.Request `json:"Most used str1"`
		Str2    dbmodels.Request `json:"Most used str2"`
		Int12   dbmodels.Request `json:"Most used int2 with int2"`
		Str12   dbmodels.Request `json:"Most used str1 with str2"`
		Strint1 dbmodels.Request `json:"Most used str1 with int1"`
		StrInt2 dbmodels.Request `json:"Most used str2 with int2"`
	}{
		*req,
		*int1,
		*int2,
		*limit,
		*str1,
		*str2,
		*str12,
		*int12,
		*strInt1,
		*strInt2,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Errorf("Could not marshel requests, err %v", err)
		return nil, err
	}
	return jsonData, nil
}
