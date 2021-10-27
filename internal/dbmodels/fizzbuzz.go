package dbmodels

import "time"

type Fizzbuzz struct {
	tableName struct{} `pg:"request"`

	ID           int64      `pg:"id" json:"-"`
	Int1         int        `pg:"int1" json:"int1"`
	Int2         int        `pg:"int2" json:"int2"`
	Limit        int        `pg:"limite" json:"limit"`
	Str1         string     `pg:"str1" json:"str1"`
	Str2         string     `pg:"str2" json:"str2"`
	CreationDate *time.Time `pg:"creation_date" json:"-"`
	LastUpdate   *time.Time `pg:"last_update" json:"-"`
}
