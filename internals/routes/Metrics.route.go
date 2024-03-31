package routes

import (
	"net/http"

	responses "github.com/victorhdchagas/nostrkeygen/internals/routes/responses"
)

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	responses.ResponseWithJSON(w, 200, Metrics{
		Requests:     0,
		Errors:       0,
		Success:      0,
		ResponseTime: 0,
		// ResponseTimes: []int64{},
	})
}

type Metrics struct {
	Requests     int64 `json:"requests"`
	Errors       int64 `json:"errors"`
	Success      int64 `json:"success"`
	ResponseTime int64 `json:"response_time"`
	// ResponseTimes []int64 `json:"response_times"`
}
