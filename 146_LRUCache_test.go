package leetcode

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test146_LRUCache(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)                   // cache is {1=1}
	lRUCache.Put(2, 2)                   // cache is {1=1, 2=2}
	assert.Equal(t, 1, lRUCache.Get(1))  // return 1
	lRUCache.Put(3, 3)                   // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	assert.Equal(t, -1, lRUCache.Get(2)) // returns -1 (not found)
	lRUCache.Put(4, 4)                   // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	assert.Equal(t, -1, lRUCache.Get(1)) // return -1 (not found)
	assert.Equal(t, 3, lRUCache.Get(3))  // return 3
	assert.Equal(t, 4, lRUCache.Get(4))  // return 4

	lRUCache = Constructor(2)
	assert.Equal(t, -1, lRUCache.Get(2)) // returns -1 (not found)
	lRUCache.Put(2, 6)                   // cache is {2=6}
	assert.Equal(t, -1, lRUCache.Get(1)) // returns -1 (not found)
	lRUCache.Put(1, 5)                   // cache is {1=5, 2=6}
	lRUCache.Put(1, 2)                   // cache is {1=2, 2=6}
	assert.Equal(t, 2, lRUCache.Get(1))  // returns 2
	assert.Equal(t, 6, lRUCache.Get(2))  // returns 6

	lRUCache = Constructor(2)
	lRUCache.Put(2, 1)                   // cache is {2=1}
	lRUCache.Put(1, 1)                   // cache is {1=1, 2=1}
	lRUCache.Put(2, 3)                   // cache is {1=1, 2=3}
	lRUCache.Put(4, 1)                   // cache is {2=3, 4=1}
	assert.Equal(t, -1, lRUCache.Get(1)) // returns -1
	assert.Equal(t, 3, lRUCache.Get(2))  // returns 3

}

type item struct {
	key   *list.Element
	value int
}

type LRUCache struct {
	maxSize int
	keys    *list.List

	cache map[int]item
}

func Constructor(capacity int) LRUCache {
	cache := make(map[int]item, capacity)

	return LRUCache{
		maxSize: capacity,
		keys:    list.New(),
		cache:   cache,
	}
}

func (c *LRUCache) Get(key int) int {
	v, ok := c.cache[key]
	if !ok {
		return -1
	}

	c.keys.MoveToFront(v.key)
	return v.value
}

func (c *LRUCache) Put(key int, value int) {
	if v, ok := c.cache[key]; ok {
		c.keys.MoveToFront(v.key)
		c.cache[key] = item{
			key:   v.key,
			value: value,
		}

		return
	}
	if c.keys.Len() >= c.maxSize {
		oldestKey := c.keys.Back()

		delete(c.cache, oldestKey.Value.(int))
		c.keys.Remove(oldestKey)
	}

	c.cache[key] = item{
		key:   c.keys.PushFront(key),
		value: value,
	}
}
