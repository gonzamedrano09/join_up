package controller

import "net/http"

type BaseController interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}
