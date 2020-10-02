# Custom-Memory-Cache

## Install
```
go get github.com/midaef/custom-memory-cache
```

## Example
```go
package tests

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
  ```
 ## Benchmark tests
 ```
goos: windows
goarch: 386
pkg: github.com/midaef/custom-memory-cache/tests
BenchmarkWrite-4         	 6519778	       188 ns/op
BenchmarkRead-4          	23077322	        53.4 ns/op
BenchmarkReadWrite-4     	 3972715	       297 ns/op
BenchmarkWriteDelete-4   	 4411705	       270 ns/op
PASS
ok  	github.com/midaef/custom-memory-cache/tests	6.288s
 ```
 
 ## License
 [MIT](https://github.com/midaef/custom-memory-cache/blob/master/LICENSE)