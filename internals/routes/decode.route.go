package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/victorhdchagas/nostrkeygen/internals/generators"
	responses "github.com/victorhdchagas/nostrkeygen/internals/routes/responses"
)

func DecodeNProfile(w http.ResponseWriter, r *http.Request) {

	nprofile := chi.URLParam(r, "nprofile")

	decodedProfile := DecodedNProfile{}
	data, err := generators.DecodeNProfile(nprofile)
	if err != nil {
		responses.ResponseWithError(w, 500, "Error decoding nprofile")
		return
	}
	decodedProfile.Pk = data.PublicKey
	decodedProfile.Relays = data.Relays
	responses.ResponseWithJSON(w, 200, decodedProfile)
}

type DecodedNProfile struct {
	Relays []string `json:"relays"`
	Pk     string   `json:"npub"`
}
