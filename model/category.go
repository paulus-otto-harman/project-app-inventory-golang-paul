package model

type Category struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	SessionId string `json:"session_id,omitempty"`
}
