package util

import (
	"fmt"
	"math/rand"
	"time"
)

/*
from Amin Shojaei @ stackoverflow.com
link: https://stackoverflow.com/a/65607935
*/
func GenerateSessionId() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 12)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[2:12]
}
