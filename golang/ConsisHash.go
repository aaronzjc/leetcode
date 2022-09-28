package golang

import (
	"hash/crc32"
	"sort"
	"strconv"

	"github.com/aaronzjc/leetcode/others"
)

// HashFunc Hash函数
type HashFunc func([]byte) uint32

// ConsistHash 一致性Hash
type ConsistHash struct {
	hash  HashFunc
	Reps  int   // 副本数
	Nodes []int // 排序后的节点
	M     map[int]string
}

// NewConsistHash 初始化
func NewConsistHash(reps int, fn HashFunc) *ConsistHash {
	if fn == nil {
		fn = crc32.ChecksumIEEE
	}
	return &ConsistHash{
		hash: fn,
		Reps: reps,
		M:    make(map[int]string),
	}
}

// Add 添加节点
func (hs *ConsistHash) Add(nodes ...string) {
	for _, node := range nodes {
		for i := 0; i < hs.Reps; i++ {
			hash := int(hs.hash([]byte(node + strconv.Itoa(i))))
			hs.Nodes = append(hs.Nodes, hash)
			hs.M[hash] = node
		}
	}

	sort.Ints(hs.Nodes)
}

// Get 根据key返回合适的节点
func (hs *ConsistHash) Get(key string) string {
	if len(hs.Nodes) == 0 {
		return ""
	}

	hash := int(hs.hash([]byte(key)))
	idx := others.BinarySearchLarge(hs.Nodes, hash)
	if idx == -1 {
		idx = 0
	}

	return hs.M[hs.Nodes[idx]]
}
