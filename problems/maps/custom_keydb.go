package maps

import (
	"errors"
	"sync"
	"time"
)

type IKeydb interface {
	Add(key string, value any, timespan time.Duration) bool
	Get(key string) (any, error)
	Delete(key string) bool
}

type entry struct {
	value     any
	expiresAt time.Time
}

type keydb struct {
	mu   sync.RWMutex
	data map[string]entry
}

func NewKeydb() IKeydb {
	return &keydb{
		data: make(map[string]entry),
	}
}

func (k *keydb) Add(key string, value any, timespan time.Duration) bool {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.data[key] = entry{
		value:     value,
		expiresAt: time.Now().Add(timespan),
	}
	return true
}

func (k *keydb) Get(key string) (any, error) {
	k.mu.RLock()
	cacheData, found := k.data[key]
	k.mu.RUnlock()

	if !found {
		return nil, errors.New("data not found")
	}

	if cacheData.expiresAt.Before(time.Now()) {
		k.Delete(key)
		return nil, errors.New("key expired or not found")
	}
	return cacheData.value, nil

}

func (k *keydb) Delete(key string) bool {
	k.mu.Lock()
	defer k.mu.Unlock()
	delete(k.data, key)
	return true
}
