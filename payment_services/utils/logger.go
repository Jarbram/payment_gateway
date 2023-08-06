package utils

import (
	"fmt"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "%v", data)
}
