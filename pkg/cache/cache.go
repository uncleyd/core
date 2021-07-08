package cache

import (
	"sync"
	"time"
)

type L interface {
	Put(key string, value interface{}, lifeTime time.Duration)
	Get(key string) interface{}
	Del(key string)
}

// 缓存对象
type CacheItem struct {
	Value     interface{}   // 实际缓存的对象
	LifeTime  time.Duration // 存活时间，上游传入
	CreatedAt time.Time     // 创建时间，和存活时间一起决定是否过期
}

// 缓存是否过期
func (item *CacheItem) Expired() bool {
	return time.Now().Sub(item.CreatedAt) > item.LifeTime
}

// 本地缓存实现类
type LocalCache struct {
	sync.RWMutex                       //继承读写锁，用于并发控制
	Items        map[string]*CacheItem // K-V存储
	GCDuration   int                   // 惰性删除，后台运行时间间隔，单位秒
}

// 新建本地缓存
func NewLocalCache(gcDuration int) L {
	localCache := &LocalCache{Items: map[string]*CacheItem{}, GCDuration: gcDuration}

	// 启动协程，定期扫描过期键，进行删除
	go localCache.GC()

	return localCache
}

// 存入对象
func (cache *LocalCache) Put(key string, value interface{}, lifeTime time.Duration) {
	cache.Lock()
	defer cache.Unlock()

	cache.Items[key] = &CacheItem{
		Value:     value,
		LifeTime:  lifeTime,
		CreatedAt: time.Now(),
	}
}

// 查询对象
func (cache *LocalCache) Get(key string) interface{} {
	cache.RLock()
	defer cache.RUnlock()

	if item, ok := cache.Items[key]; ok {
		if !item.Expired() {
			return item
		} else {
			// 键已过期，直接删除
			// 需要注意的是，这里不能调用cache.Del()方法，因为go的读写锁是不支持锁升级的，会发生死锁
			delete(cache.Items, key)
		}
	}

	return nil
}

// 删除缓存
func (cache *LocalCache) Del(key string) {
	cache.Lock()
	defer cache.Unlock()

	if _, ok := cache.Items[key]; ok {
		delete(cache.Items, key)
	}
}

// 异步执行，扫描过期键并删除
func (cache *LocalCache) GC() {
	for {
		select {
		case <-time.After(time.Duration(cache.GCDuration) * time.Second):
			keysToExpire := []string{}

			cache.RLock()
			for key, item := range cache.Items {
				if item.Expired() {
					keysToExpire = append(keysToExpire, key)
				}
			}
			cache.RUnlock()

			for _, keyToExpire := range keysToExpire {
				cache.Del(keyToExpire)
			}
		}
	}
}
