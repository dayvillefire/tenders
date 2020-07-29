package common

import (
	"math/rand"
	"strconv"
)

// ShortCode generates a new short code of specified length
func ShortCode(length int) string {
	result := ""
	for i := 0; i < length; i++ {
		r := rand.Intn(35)
		if r < 10 {
			result += strconv.Itoa(r)
		} else {
			result += string(r + 87)
		}
	}
	return result
}