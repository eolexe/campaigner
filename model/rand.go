package model

import (
	"math/rand"
	"time"
)

var (
	Randomizer randomizer = rand.New(rand.NewSource(time.Now().Unix()))
)

type randomizer interface {
	Int63n(n int64) int64
}
