package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Controller controls the handles the response to an HTTP request
type Controller struct {
	handler func(params map[string]string) []byte
}

func (c Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	vals := r.Form

	for k, v := range vals {
		if len(v) == 0 {
			continue
		}
		vars[k] = v[0]
	}

	data := c.handler(vars)
	w.Write(data)
}

// NewController creates a Controller with passed in param func as its callback
func NewController(h func(params map[string]string) []byte) *Controller {
	return &Controller{
		handler: h,
	}
}
