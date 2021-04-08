package golang

import (
	"github.com/aaronzjc/leetcode/others"
	"hash/crc32"
	"sort"
	"strconv"
)

// 一致性hash
type HashFunc func([]byte) uint32

type ConsisHash struct {
	hash  HashFunc
	Reps  int   // 副本数
	Nodes []int // 排序后的节点
	M     map[int]string
}

func NewConsistHash(reps int, fn HashFunc) *ConsisHash {
	if fn == nil {
		fn = crc32.ChecksumIEEE
	}
	return &ConsisHash{
		hash: fn,
		Reps: reps,
		M:    make(map[int]string),
	}
}

// Add 添加节点
func (hs *ConsisHash) Add(nodes ...string) {
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
func (hs *ConsisHash) Get(key string) string {
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
