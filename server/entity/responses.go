package entity

type Resp struct {
	Success bool   `json:"success"`
	ErrNo   int    `json:"errno"`
	Errmsg  string `json:"errmsg"`
}

type RespPutRecord struct {
	Resp
	Data Record `json:"data"`
}

type RespGetRecord struct {
	Resp
	Data Record `json:"data"`
}
