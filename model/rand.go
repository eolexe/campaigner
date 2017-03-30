package model

import (
	"math/rand"
	"sync"
	"time"
)

var (
	Rnd      randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))
	RndMutex            = &sync.Mutex{}
)

type randomizer interface {
	Int63n(n int64) int64
}
