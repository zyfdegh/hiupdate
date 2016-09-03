package entity

// Resp is the common struct of http response body
type Resp struct {
	Success bool   `json:"success"`
	ErrNo   int    `json:"errno"`
	Errmsg  string `json:"errmsg"`
}

// RespPutRecord is the response to PUT /record
type RespPutRecord struct {
	Resp
	Data Record `json:"data"`
}

// RespGetRecord is the response to GET /record
type RespGetRecord struct {
	Resp
	Data Record `json:"data"`
}
