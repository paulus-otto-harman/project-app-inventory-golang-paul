package model

type Location struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	SessionId string `json:"session_id,omitempty"`
}
