package models

type ResponseDefaultModel struct {
	IsSuccess    bool        `json:"isSuccess"`
	Data         interface{} `json:"data"`
	Error        interface{} `json:"error,omitempty"`
	ErrorMessage string      `json:"errorMessage,omitempty"`
}
