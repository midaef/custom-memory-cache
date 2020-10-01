package cache

import "time"

type GC struct {
	Cache *Cache
}

func NewGC(cache *Cache) *GC {
	return &GC{
		Cache: cache,
	}
}

func (gc *GC) Start() {
	go gc.garbageCollector()
}

func (gc *GC) garbageCollector() {
	for {
		<-time.After(gc.Cache.GCInterval)
		keys := gc.searchKeys()
		if len(keys) != 0 {
			gc.Cache.DataMutex.Lock()
			for _, key := range keys {
				delete(gc.Cache.Data, key)
			}
			gc.Cache.DataMutex.Unlock()
		}
	}
}

func (gc *GC) searchKeys() []string {
	var keys []string
	gc.Cache.DataMutex.Lock()
	defer gc.Cache.DataMutex.Unlock()
	if gc.Cache.Data == nil {
		return keys
	}
	for k, value := range gc.Cache.Data {
		if time.Now().UnixNano() > value.LifeTime && value.LifeTime > 0 {
			value.Alive = false
			keys = append(keys, k)
		}
	}
	return keys
}
