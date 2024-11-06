package model

type DataPage struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Page       int         `json:"page,omitempty"`
	Limit      int         `json:"limit,omitempty"`
	TotalItems int         `json:"total_items,omitempty"`
	TotalPages float64     `json:"total_pages,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
