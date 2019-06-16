package helper

import (
	"encoding/json"
	"net/http"
)

// ResponseCallback send response to client
func ResponseCallback(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
