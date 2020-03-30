package model

type Result struct {
	Rows []Row `json:"rows"`
	Affects int64 `json:"affects"`
	Error error `json:"error"`
}
