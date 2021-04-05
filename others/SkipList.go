/**
 * 跳跃表
 * 一句话总结跳跃表就是，基于链表实现二分搜索，理论性能能达到log(n)。
 * 之所以是理论，是因为它的层数不是固定，而是随机。为什么是随机？因为没办法预料元素的分布，如果每次插入元素都动态调整层数，那会相当复杂。
 */

package others

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const MAX_LEVEL = 8

var (
	MAX_INT = 1<<32 - 1
	MIN_INT = ^MAX_INT
)

type SkipListNode struct {
	val int

	level []*SkipListNode
}

type SkipList struct {
	head, tail *SkipListNode
	len        int
}

func NewSkipList() *SkipList {
	list := &SkipList{
		head: &SkipListNode{
			val:   MIN_INT,
			level: make([]*SkipListNode, MAX_LEVEL, MAX_LEVEL),
		},
		tail: &SkipListNode{
			val:   MAX_INT,
			level: make([]*SkipListNode, MAX_LEVEL, MAX_LEVEL),
		},
	}
	for i := 0; i < MAX_LEVEL; i++ {
		list.head.level[i] = list.tail
	}
	return list
}

func RandomLevel() int {
	level := 1
	for {
		rand.Seed(time.Now().UnixNano())
		if rand.Int()%4 == 1 || level >= MAX_LEVEL {
			break
		}
		level++
	}

	return level
}

func (sl *SkipList) Add(val int) *SkipListNode {
	fmt.Println(val)
	var update [MAX_LEVEL]*SkipListNode
	p := sl.head
	for i := MAX_LEVEL - 1; i >= 0; i-- {
		for i < len(p.level) && p.level[i].val < val {
			p = p.level[i]
		}
		update[i] = p
	}
	// 不可能超过最大高度
	level := RandomLevel()
	node := &SkipListNode{
		val:   val,
		level: make([]*SkipListNode, level, level),
	}
	for i := level - 1; i >= 0; i-- {
		p = update[i]
		node.level[i] = p.level[i]
		p.level[i] = node
	}

	return node
}

func (sl *SkipList) Get(val int) (*SkipListNode, error) {
	p := sl.head
	for i := MAX_LEVEL - 1; i >= 0; i-- {
		for i < len(p.level) && p.level[i].val <= val {
			p = p.level[i]
			if p.val == val {
				return p, nil
			}
		}
	}
	return nil, errors.New("not found")
}

func (sl *SkipList) Dump() {
	p := sl.head
	for p != nil {
		for k, v := range p.level {
			if v != nil {
				fmt.Printf("[%d] = %d ", k, v.val)
			}
		}
		fmt.Printf("%d\n", p.val)
		p = p.level[0]
	}
}
