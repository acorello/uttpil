package uttpil

import (
	"fmt"
	"net/http"
)

func RespondErrMethodNotImplemented(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("method %q not allowed", r.Method)
	http.Error(w, msg, http.StatusMethodNotAllowed)
}
