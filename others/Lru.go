package others

import (
	"fmt"
	"strconv"
)

// LRU - 最近最少使用算法

type LRUKey int
type LRUVal int

type LRUNode struct {
	key  LRUKey
	val  LRUVal
	pre  *LRUNode
	next *LRUNode
}

type LRUCache struct {
	size  int
	start *LRUNode
	end   *LRUNode
	hm    map[LRUKey]*LRUNode
}

type LRUDoer interface {
	Set(key LRUKey, val LRUVal)
	Get(key LRUKey) (LRUVal, bool)
	Dump() string
}

func NewLRUCache(size int) LRUDoer {
	return &LRUCache{
		size: size,
		hm:   make(map[LRUKey]*LRUNode),
	}
}

func (cache *LRUCache) Set(key LRUKey, val LRUVal) {
	e, ok := cache.hm[key]
	if !ok {
		// 如果元素不存在
		// 1. 初始化一个节点
		// 2. 判断队列是否满了，如果满了，则删除一个元素
		// 3. 插入元素到队列头部
		node := &LRUNode{
			key: key,
			val: val,
		}
		if len(cache.hm) == cache.size {
			cache.removeNode(cache.end)
		}
		cache.addNodeAtTop(node)
		return
	}

	// 如果元素存在
	// 从原来的位置移除，挪到顶部
	cache.removeNode(e)
	cache.addNodeAtTop(e)
}

func (cache *LRUCache) Get(key LRUKey) (LRUVal, bool) {
	e, ok := cache.hm[key]
	if ok {
		cache.removeNode(e)
		cache.addNodeAtTop(e)
		return e.val, true
	}

	return LRUVal(0), false
}

func (cache *LRUCache) Dump() string {
	fmt.Println(cache.hm)
	node := cache.start
	res := ""
	for node != nil {
		res += strconv.Itoa(int(node.val))
		if node.next != nil {
			res += ","
		}
		node = node.next
	}

	return res
}

func (cache *LRUCache) addNodeAtTop(node *LRUNode) {
	cache.hm[node.key] = node
	if cache.start == nil {
		cache.end = node
		cache.start = node
		return
	}

	node.next = cache.start
	cache.start.pre = node
	cache.start = node
}

func (cache *LRUCache) removeNode(node *LRUNode) {
	delete(cache.hm, node.key)

	if node.pre != nil {
		node.pre.next = node.next
	} else {
		cache.start = node.next
	}

	if node.next != nil {
		node.next.pre = node.pre
	} else {
		cache.end = node.pre
	}

	node.pre = nil
	node.next = nil
}
