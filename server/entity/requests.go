package entity

// Req to /update
type ReqUpdate struct {
	Name  string `json:"name"`
	Done  string `json:"done"`
	Todo  string `json:"todo"`
	Issue string `json:"issue"`
}
