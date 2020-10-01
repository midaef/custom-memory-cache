package main

import (
	"log"

	"github.com/midaef/custom-memory-cache/cache"
)

func main() {
	// Create new cache function takes two arguments.
	// First cache size and second garbage collerctor interval.
	// gcInterval it's time between clearing the cache
	// if set gcInterval - 0 then garbage collerctor won't work
	cache := cache.NewCache(16, 0)

	// Write data with key to cache
	cache.Write("token", "1234", 0)

	// Read data from cache by key
	data, err := cache.Read("token")
	if err != nil {
		log.Println(err)
	}
	log.Println(data)

	// Read all data from cache
	mapData, err := cache.ReadAll()
	if err != nil {
		log.Println(err)
	}
	log.Println(mapData)

	// Delete element from cache by key
	err = cache.Delete("token")
	if err != nil {
		log.Println(err)
	}

	// Delete all elements from cache
	cache.Write("token", "1234", 0)
	err = cache.DeleteAll()
	if err != nil {
		log.Println(err)
	}
}
