package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func Int(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
