package controller

import "net/http"

type Action interface {
	Action(w http.ResponseWriter, r *http.Request)
}
