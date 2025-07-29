package store

import (
	"sync"
	"time"
)

type Sets struct {
	Value    string
	ExpireAt time.Time
}

var (
	SETs  = map[string]Sets{}
	SetMu = sync.RWMutex{}

	HSETs  = map[string]map[string]string{}
	HsetMu = sync.RWMutex{}

	expQ   = &ExpiryHeap{}
	heapmu = sync.Mutex{}
)
