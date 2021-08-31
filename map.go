package syncmap

import (
	"sync"
)

// Map is a thread-safe map mapping from typed key to typed value.
type Map[K, V comparable] struct {
	m    map[K]V
	lock sync.RWMutex
}

// New returns a new Map.
func New[K, V comparable](m map[K]V) *Map[K, V] {
	if m == nil {
		m = make(map[K]V)
	}
	return &Map[K, V]{
		m: m,
	}
}

// Get returns a value of k, it returns nil if not found.
func (s *Map[K, V]) Get(k K) (V, bool) {
	s.lock.RLock()
	v, ok := s.m[k]
	s.lock.RUnlock()
	return v, ok
}

// Set sets value v to key k in the map.
func (s *Map[K, V]) Set(k K, v V) {
	s.lock.Lock()
	s.m[k] = v
	s.lock.Unlock()
}

// Update updates value v to key k, returns false if k not found.
func (s *Map[K, V]) Update(k K, v V) bool {
	s.lock.Lock()
	_, ok := s.m[k]
	if !ok {
		s.lock.Unlock()
		return false
	}
	s.m[k] = v
	s.lock.Unlock()
	return true
}

// Delete deletes a key in the map.
func (s *Map[K, V]) Delete(k K) {
	s.lock.Lock()
	delete(s.m, k)
	s.lock.Unlock()
}

// Dup duplicates the map to a new map.
func (s *Map[K, V]) Dup() *Map[K, V] {
	newMap := New[K, V](nil)
	s.lock.RLock()
	for k, v := range s.m {
		newMap.m[k] = v
	}
	s.lock.RUnlock()
	return newMap
}
