package state

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"strconv"
)

func getState(nonce []byte, secret string) uint64 {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(nonce))
	mac := h.Sum(nonce)
	var state uint64
	for i := 0; i < 8; i++ {
		state = state<<8 | uint64(mac[i])
	}
	return state
}

// Generate initializes a new state using random and secret string
func Generate(secret string) string {
	nonce := make([]byte, 4)
	_, err := rand.Read(nonce)
	if err != nil {
		return ""
	}
	state := getState(nonce, secret)
	return strconv.FormatUint(state, 10)
}

// Valid checks state
func Valid(stateString, secret string) bool {
	state, err := strconv.ParseUint(stateString, 10, 64)
	if err != nil {
		return false
	}
	nonce := make([]byte, 4)
	j := state >> 32
	for i := 3; i >= 0; i-- {
		nonce[i] = byte(0xff & j)
		j = j >> 8
	}
	check := getState(nonce, secret)
	return state == check
}
