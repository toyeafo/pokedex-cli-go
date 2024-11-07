package pokecache

import "time"

type Cache struct {
	entry map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
