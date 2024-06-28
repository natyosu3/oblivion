package general

import (
	"math/rand"

	"github.com/google/uuid"
)

const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func MakeRandomId() string {
	randomVal := uuid.New()

	return randomVal.String()
}

// n桁のランダムな文字列を生成する
func MakeRandomStringId(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(b)
}
