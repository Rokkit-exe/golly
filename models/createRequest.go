package models

type CreateRequest struct {
	Name   string `json:"name"`
	From   string `json:"from"`
	System string `json:"system"`
}
