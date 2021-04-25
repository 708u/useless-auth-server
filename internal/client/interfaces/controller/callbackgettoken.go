package controller

import "net/http"

type CallbackGetToken struct {
}

func NewCallbackGetToken() *CallbackGetToken {
	return &CallbackGetToken{}
}

func (c *CallbackGetToken) Action(w http.ResponseWriter, r *http.Request) {
	//
}
