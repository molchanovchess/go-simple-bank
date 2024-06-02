package util

import (
	"math/rand"
	"time"
)

// Initialize a new rand.Rand instance with a seed
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomInt generates a random integer between min and max (inclusive)
func RandomInt(min, max int64) int64 {
	if min > max {
		min, max = max, min
	}
	return min + rng.Int63n(max-min+1)
}

// RandomString generates a random string of a specified length
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rng.Intn(len(charset))]
	}
	return string(result)
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	return currencies[rng.Intn(len(currencies))]
}
