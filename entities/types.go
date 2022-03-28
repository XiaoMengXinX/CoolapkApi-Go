package ent

import "net/http"

type RawData struct {
	Header   http.Header `json:"-"`
	Response string      `json:"-"`
}

type ErrorMsg struct {
	Status        int    `json:"status"`
	Error         int    `json:"error"`
	Message       string `json:"message"`
	MessageStatus int    `json:"messageStatus"`
}
