package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Token struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

func tokenMaker() string {
	x := time.Now().UnixNano()
	sum := strconv.FormatInt(x, 10)
	x2 := sha256.Sum256([]byte(sum))

	return fmt.Sprintf("%x", x2)
}

func Tokenapi(w http.ResponseWriter, r *http.Request) {
	var tokenr Token
	tokenr.Id = "1"
	tokenr.Token = tokenMaker()

	json.NewEncoder(w).Encode(tokenr)

}
