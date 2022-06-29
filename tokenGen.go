package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

func tokenMaker() string {
	x := time.Now().UnixNano()
	sum := strconv.FormatInt(x, 10)
	x2 := sha256.Sum256([]byte(sum))

	return fmt.Sprintf("%x", x2)
}
