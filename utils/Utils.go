package utils

import (
	"net/http"
	"encoding/json"
)

var AccessKey string

func WriteJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("\n", "")
	encoder.Encode(data)
}

func IsAccessTokenValid(r *http.Request) bool {
	if r.PostFormValue("access_token") == AccessKey {
		return true
	}

	return false
}