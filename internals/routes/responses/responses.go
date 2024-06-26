package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %v", msg)

	}

	type errResponse struct {
		Error string `json:"error"`
	}

	ResponseWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
