package random

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	options := "Ñ765ASDFGHJKLPO1234IUYTREWQZXCVBNM098"
	rand.Seed(time.Now().UnixNano())
	rs := make([]byte, length)
	for i := range rs {
		rs[i] = options[rand.Intn(len(options))]
	}
	return string(rs)
}
