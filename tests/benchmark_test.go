package tests

import (
	"testing"

	"github.com/midaef/custom-memory-cache/cache"
)

func BenchmarkWrite(b *testing.B) {
	cache := cache.NewCache(1024, 1)
	for i := 0; i < b.N; i++ {
		cache.Write("test", "test", 1)
	}
}

func BenchmarkRead(b *testing.B) {
	cache := cache.NewCache(1024, 1)
	cache.Write("test", "test", 1)
	for i := 0; i < b.N; i++ {
		cache.Read("test")
	}
}

func BenchmarkReadWrite(b *testing.B) {
	cache := cache.NewCache(1024, 1)
	for i := 0; i < b.N; i++ {
		cache.Write("test", "test", 1)
		cache.Read("test")
	}
}

func BenchmarkWriteDelete(b *testing.B) {
	cache := cache.NewCache(1024, 1)
	for i := 0; i < b.N; i++ {
		cache.Write("test", "test", 1)
		cache.Delete("test")
	}
}
