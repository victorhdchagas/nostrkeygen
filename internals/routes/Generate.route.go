package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/victorhdchagas/nostrkeygen/internals/generators"
	responses "github.com/victorhdchagas/nostrkeygen/internals/routes/responses"
)

func GetKeys(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Words []string `json:"words"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responses.ResponseWithError(w, 400, fmt.Sprintf("Error parsing Json:%v", err))
	}
	nostrKeys := NostrKeys{}
	nostrKeys.Sk = generators.GenerateKey(params.Words)
	Pk, err := generators.GeneratePublicKey(nostrKeys.Sk)

	if err != nil {
		responses.ResponseWithError(w, 400, fmt.Sprintf("Error generating public key:%v", err))
	}
	nostrKeys.Pk = Pk
	Nsec, err := generators.GenerateNsec(nostrKeys.Sk)
	if err != nil {
		responses.ResponseWithError(w, 400, fmt.Sprintf("Error generating nsec:%v", err))
	}
	nostrKeys.Nsec = Nsec
	npub, err := generators.GenerateNPub(nostrKeys.Pk)
	if err != nil {
		responses.ResponseWithError(w, 400, fmt.Sprintf("Error generating npub:%v", err))
	}
	nostrKeys.Npub = npub
	responses.ResponseWithJSON(w, 200, nostrKeys)
}

type NostrKeys struct {
	Sk   string `json:"sk"`
	Pk   string `json:"pk"`
	Nsec string `json:"nsec"`
	Npub string `json:"npub"`
}
