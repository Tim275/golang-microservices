// common/json.go
package common

import (
	"encoding/json"
	"net/http"
)

// ReadJson liest JSON-Daten aus der Anfrage und decodiert sie in das angegebene Ziel.
func ReadJson(r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(target)
}

// WriteError schreibt eine JSON-Fehlermeldung mit dem gegebenen Statuscode.
func WriteError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

// WriteJSON schreibt eine JSON-Antwort mit dem gegebenen Statuscode.
func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
