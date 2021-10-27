package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"

	"lbc/fizzbuzz/internal/dbmodels"
)

func AggReqCount(db *pg.DB) (*dbmodels.Request, error) {
	var requests []dbmodels.Request
	err := db.Model((*dbmodels.Fizzbuzz)(nil)).
		Column("int1", "int2", "limite", "str1", "str2").
		ColumnExpr("count(*) AS count").
		Group("int1", "int2", "limite", "str1", "str2").
		Limit(1).
		Order("count DESC").
		Select(&requests)
	if err != nil {
		logrus.Errorf("Could not query requests, err %v", err)
		return nil, err
	}
	return &requests[0], nil
}

func AggFieldCount(db *pg.DB, keys ...string) (*dbmodels.Request, error) {
	var aggs []dbmodels.Request
	err := db.Model((*dbmodels.Fizzbuzz)(nil)).
		Column(keys...).
		ColumnExpr("count(*) AS count").
		Group(keys...).
		Limit(1).
		Order("count DESC").
		Select(&aggs)
	if err != nil {
		logrus.Errorf("Could not query requests, err %v", err)
		return nil, err
	}
	return &aggs[0], nil
}
