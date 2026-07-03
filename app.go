package main

import (
	"context"
	"sync"
)

type JobUpdate struct {
	Status    string  `json:"status"`
	Progress  float64 `json:"progress"`
	Processed int     `json:"processed"`
	Total     int     `json:"total"`
	Error     string  `json:"error,omitempty"`
}

type jobBroadcaster struct {
	mu          sync.Mutex
	subscribers map[int64]chan JobUpdate
	nextID      int64
	current     JobUpdate
}

func newJobBroadcaster() *jobBroadcaster {
	return &jobBroadcaster{
		subscribers: make(map[int64]chan JobUpdate),
	}
}

func (b *jobBroadcaster) subscribe() (chan JobUpdate, func()) {
	b.mu.Lock()
	defer b.mu.Unlock()
	ch := make(chan JobUpdate, 64)
	id := b.nextID
	b.nextID++
	b.subscribers[id] = ch
	ch <- b.current
	return ch, func() {
		b.mu.Lock()
		delete(b.subscribers, id)
		b.mu.Unlock()
	}
}

func (b *jobBroadcaster) broadcast(u JobUpdate) {
	b.mu.Lock()
	b.current = u
	for _, ch := range b.subscribers {
		select {
		case ch <- u:
		default:
		}
	}
	b.mu.Unlock()
}

type PendingJob struct {
	ID          string
	Params      FetchParams
	FilePath    string
	FileName    string
	Ctx         context.Context
	Cancel      context.CancelFunc
	Broadcaster *jobBroadcaster
}
