package store

import (
	"sync"
)

type Store struct {
	mu    sync.RWMutex
	store map[string]any
}

func NewStore() *Store {
	return &Store{store: make(map[string]any, 20)}
}

func (this *Store) Add(key string, value any) {
	this.mu.Lock()
	this.store[key] = value
	this.mu.Unlock()
}

func (this *Store) Remove(key string) bool {
	startLen := len(this.store)

	this.mu.Lock()
	delete(this.store, key)
	this.mu.Unlock()

	newLen := len(this.store)

	return startLen != newLen
}

func (this *Store) Get(key string) any {
	this.mu.RLock()
	val := this.store[key]
	this.mu.RUnlock()

	return val
}
