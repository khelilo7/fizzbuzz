package dbmodels

type Request struct {
	Int1   int    `json:"int1,omitempty"`
	Int2   int    `json:"int2,omitempty"`
	Limite int    `json:"limit,omitempty"`
	Str1   string `json:"str1,omitempty"`
	Str2   string `json:"str2,omitempty"`
	Count  int    `json:"count,omitempty"`
}
