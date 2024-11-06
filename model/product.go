package model

type Product struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CategoryId int    `json:"category_id"`
	LocationId int    `json:"location_id"`
	Balance    int    `json:"balance"`
	SessionId  string `json:"session_id,omitempty"`
}
