package cache

import (
	"log"
)

type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	GetStat() Stat
	NewScanner() Scanner
}

func New(typ string, ttl int) Cache {
	var c Cache
	switch typ {
	case "inmemory":
		c = newInMemoryCache(ttl)
	case "rocksdb":
		c = newRocksdbCache(ttl)
	default:
		panic("unkown cache type: " + typ)
	}
	if c == nil {
		panic("unkown cache type: " + typ)
	}
	log.Println(typ, "ready to serve")
	return c
}
