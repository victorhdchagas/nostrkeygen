package generators

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

func wordsToHex(words []string) string {
	seedPhrase := strings.Join(words, "")
	hash := sha256.Sum256([]byte(seedPhrase))
	hexString := hex.EncodeToString(hash[:])

	return hexString
}

func GenerateKey(words []string) string {
	hexString := wordsToHex(words)
	return hexString
}

func GenerateNsec(sk string) (string, error) {
	return nip19.EncodePrivateKey(sk)
}

func GeneratePublicKey(sk string) (string, error) {
	return nostr.GetPublicKey(sk)
}

func GenerateNPub(pk string) (string, error) {
	return nip19.EncodePublicKey(pk)
}

func DecodeNProfile(npub string) (nostr.ProfilePointer, error) {
	_, data, err := nip19.Decode(npub) //"nprofile1qy88wumn8ghj7mn0wvhxcmmv9uq3zamnwvaz7tmwdaehgu3wwa5kuef0qyt8wumn8ghj7etyv4hzumn0wd68ytnvv9hxgtcqyqduwzspfzelx9k6x0lrez0j8cl8rtz0lxvqylk8z2ustnfy76jpzeeu00q"
	if err != nil {
		return nostr.ProfilePointer{}, err
	}
	pp, ok := data.(nostr.ProfilePointer)
	if !ok {
		return nostr.ProfilePointer{}, errors.New("invalid profile pointer")
	}
	return pp, nil
}
