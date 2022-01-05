package cmd

import (
	"math/rand"
	"strings"
	"time"
)

type CLIFlags struct {
	Target          string
	Targets         []string
	OSTargets       []string
	LocationTargets []string
	UserTargets     []string
	Query           string
}

func interfaceStrings(v []string) []interface{} {
	a := make([]interface{}, len(v))
	for i := 0; i < len(v); i++ {
		a[i] = v[i]
	}
	return a
}

// RandomString returns a random alphanumeric string.
func RandomString(n int) string {
	const (
		letterBytes   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	src := rand.NewSource(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return sb.String()
}
