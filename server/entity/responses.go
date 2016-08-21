package entity

type Resp struct {
	Success bool   `json:"success"`
	ErrNo   int    `json:"errno"`
	Errmsg  string `json:"errmsg"`
}

type RespPutUpdate struct {
	Resp
	Data Record `json:"data"`
}

type RespGetUpdate struct {
	Resp
	Data Record `json:"data"`
}
