package entity

type Resp struct {
	Success bool   `json:"success"`
	ErrNo   int    `json:"errno"`
	Errmsg  string `json:"errmsg"`
}

type RespUpdate struct {
	Resp
	Data ReqUpdate `json:"data"`
}
