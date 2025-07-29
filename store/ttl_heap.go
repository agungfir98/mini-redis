package store

import (
	"container/heap"
	"time"
)

func init() {
	heap.Init(expQ)
	go ExpireCleaner()
}

type ExpiryItem struct {
	key      string
	expireAt time.Time
}

type ExpiryHeap []ExpiryItem

func (h ExpiryHeap) Len() int           { return len(h) }
func (h ExpiryHeap) Less(i, j int) bool { return h[i].expireAt.Before(h[j].expireAt) }
func (h ExpiryHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ExpiryHeap) Push(x any) {
	*h = append(*h, x.(ExpiryItem))
}
func (h *ExpiryHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func ExpireCleaner() {
	for range time.Tick(time.Second) {
		now := time.Now()
		heapmu.Lock()
		for expQ.Len() > 0 {
			item := (*expQ)[0]
			if now.Before(item.expireAt) {
				break
			}
			heap.Pop(expQ)
			heapmu.Unlock()

			SetMu.Lock()
			entry, ok := SETs[item.key]
			if ok && entry.ExpireAt.Equal(item.expireAt) {
				delete(SETs, item.key)
			}
			SetMu.Unlock()
			heapmu.Lock()
		}
		heapmu.Unlock()
	}
}
