package model

type Row struct {
	Data map[string]interface{} `json:"data"`
	Type int `json:"type"`
}