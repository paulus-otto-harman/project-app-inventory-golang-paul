package model

type Transaction struct {
	Id          int    `json:"id,omitempty"`
	Type        string `json:"type"`
	ProductId   int    `json:"product_id"`
	Quantity    int    `json:"quantity"`
	Note        string `json:"note"`
	SessionId   string `json:"session_id"`
	DateCreated string `json:"created_at"`
}
