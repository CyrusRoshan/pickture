package cache

import (
	"image"
	"sync"
)

// Image cache uses an array structure for O[1] indexing.
// This requires knowing the size of the cache beforehand.
type ImageCache struct {
	images []*image.Image
	locks  []sync.Mutex

	ImageCacheProps
}

// We need an array of all of the paths for the images as well as a lookup
// function.
type ImageCacheProps struct {
	Paths    []string
	LoadFunc func(string) image.Image

	// Number of images to preload ahead of i when calling Get(i)
	PreloadCount int
}

func NewImageCache(props ImageCacheProps) *ImageCache {
	length := len(props.Paths)
	cache := ImageCache{
		images:          make([]*image.Image, length),
		locks:           make([]sync.Mutex, length),
		ImageCacheProps: props,
	}

	// go cache.Preload(1) // Preload first item
	return &cache
}

func (c *ImageCache) Get(i int) *image.Image {
	c.locks[i].Lock()
	img := c.images[i]
	if img == nil {
		loadedImg := c.LoadFunc(c.Paths[i])
		img = &loadedImg
		c.images[i] = img
	}
	c.locks[i].Unlock()

	// Preload ahead
	go func(i int) {
		for j := 1; j < c.PreloadCount; j++ {
			index := i + j
			c.Preload(index)
		}
	}(i)

	return img
}

func (c *ImageCache) Preload(i int) {
	if c.images[i] == nil {
		loadedImg := c.LoadFunc(c.Paths[i])
		c.images[i] = &loadedImg
	}
}
