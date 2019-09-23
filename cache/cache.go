package cache

import (
	"sync"

	"github.com/gotk3/gotk3/gdk"
)

// FileArrayCache uses an array structure for O[1] indexing.
// This requires knowing the size of the cache beforehand.
type FileArrayCache struct {
	items []*gdk.Pixbuf
	locks []sync.Mutex

	FileArrayCacheProps
}

// We need an array of all of the paths as well as a lookup function to perform ETL.
type FileArrayCacheProps struct {
	Paths    []string
	LoadFunc func(string) *gdk.Pixbuf

	// Number of files to preload ahead of i when calling Get(i)
	PreloadCount int
}

func NewFileArrayCache(props FileArrayCacheProps) *FileArrayCache {
	length := len(props.Paths)
	cache := FileArrayCache{
		items:               make([]*gdk.Pixbuf, length),
		locks:               make([]sync.Mutex, length),
		FileArrayCacheProps: props,
	}

	return &cache
}

func (c *FileArrayCache) Get(i int) *gdk.Pixbuf {
	if c.outOfBounds(i) {
		return nil
	}
	c.locks[i].Lock()

	item := c.items[i]
	if item == nil {
		c.items[i] = c.LoadFunc(c.Paths[i])
		item = c.items[i]
	}

	c.locks[i].Unlock()

	// Preload ahead
	go func(i int) {
		for j := 1; j < c.PreloadCount; j++ {
			index := i + j
			c.Preload(index)
		}
	}(i)

	return item
}

func (c *FileArrayCache) Preload(i int) {
	if c.outOfBounds(i) {
		return
	}
	c.locks[i].Lock()
	if c.items[i] == nil {
		c.items[i] = c.LoadFunc(c.Paths[i])
	}
	c.locks[i].Unlock()
}

func (c *FileArrayCache) outOfBounds(i int) bool {
	return i < 0 || i > len(c.items)-1
}
