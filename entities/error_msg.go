package ent

type ErrorMsg struct {
	Status        int    `json:"status"`
	Error         int    `json:"error"`
	Message       string `json:"message"`
	MessageStatus int    `json:"messageStatus"`
}
