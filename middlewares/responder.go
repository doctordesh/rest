package middlewares

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Responder interface {
	SendError(int, string)
	SendJson(interface{}) error
}

type standardResponder struct {
	ResponseWriter http.ResponseWriter
}

func NewResponder(w http.ResponseWriter) Responder {
	s := new(standardResponder)
	s.ResponseWriter = w
	return s
}

func (r *standardResponder) SendError(code int, reason string) {
	r.ResponseWriter.WriteHeader(code)
	json.NewEncoder(r.ResponseWriter).Encode(map[string]string{
		"status":  http.StatusText(code),
		"code":    strconv.Itoa(code),
		"message": reason,
	})
}

func (r *standardResponder) SendJson(obj interface{}) error {
	return json.NewEncoder(r.ResponseWriter).Encode(obj)
}
