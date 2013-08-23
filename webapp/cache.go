package main

import (
        "sync"
        "time"
)

type Entry struct {
        Expire  time.Time
        Content []byte
}

type Cache struct {
        cache map[string]*Entry
        mutex *sync.RWMutex
}

var cache = NewCache()

//var cache  = map[string] *Entry {}

func NewCache() (c *Cache) {
        var mapcache = map[string]*Entry{}
        c = &Cache{cache: mapcache, mutex: new(sync.RWMutex)}
        return
}

func (c *Cache) getCachedPage(url string) []byte {
        c.mutex.RLock()
        defer c.mutex.RUnlock()
        if c.cache[url] == nil {
                return nil
        }
        if c.cache[url].Expire.Before(time.Now()) {
                return nil
        }
        return c.cache[url].Content
}

func (c *Cache) putCachedPage(url string, data []byte) {
        dur := 4 * time.Hour
        c.mutex.Lock()
        c.cache[url] = &Entry{Expire: time.Now().Add(dur), Content: data}
        c.mutex.Unlock()
        return
}
