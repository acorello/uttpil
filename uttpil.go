package uttpil

import (
	"log"
	"net/http"
)

type ForMethod struct {
	CONNECT,
	DELETE,
	GET,
	HEAD,
	OPTIONS,
	PATCH,
	POST,
	PUT,
	TRACE http.HandlerFunc
}

func (h ForMethod) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if h.GET != nil {
			h.GET(w, r)
		}
	case http.MethodPost:
		if h.POST != nil {
			h.POST(w, r)
		}
	case http.MethodPut:
		if h.PUT != nil {
			h.PUT(w, r)
		}
	case http.MethodDelete:
		if h.DELETE != nil {
			h.DELETE(w, r)
		}
	case http.MethodHead:
		if h.HEAD != nil {
			h.HEAD(w, r)
		}
	case http.MethodOptions:
		if h.OPTIONS != nil {
			h.OPTIONS(w, r)
		}
	case http.MethodPatch:
		if h.PATCH != nil {
			h.PATCH(w, r)
		}
	case http.MethodTrace:
		if h.TRACE != nil {
			h.TRACE(w, r)
		}
	case http.MethodConnect:
		if h.CONNECT != nil {
			h.CONNECT(w, r)
		}
	default:
		log.Printf("no handler for method %q", r.Method)
		http.NotFound(w, r)
	}
}
