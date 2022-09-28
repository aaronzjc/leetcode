/**
 * 无锁队列
 * 看样子不是无锁的。链表实现的有锁队列。
 */

package golang

import (
	"errors"
	"sync"
)

// NonBlockQueueItem 队列元素
type NonBlockQueueItem struct {
	Val  int
	next *NonBlockQueueItem
}

// NonBlockQueue 无锁队列
type NonBlockQueue struct {
	pushLock *sync.Mutex
	popLock  *sync.Mutex

	head *NonBlockQueueItem
	tail *NonBlockQueueItem
}

// NewNonBlockQueue 初始化
func NewNonBlockQueue() *NonBlockQueue {
	return &NonBlockQueue{
		pushLock: &sync.Mutex{},
		popLock:  &sync.Mutex{},
	}
}

// Push 添加
func (q *NonBlockQueue) Push(item int) {
	node := &NonBlockQueueItem{
		Val:  item,
		next: nil,
	}

	// TODO: using cas instead of lock
	q.pushLock.Lock()
	defer q.pushLock.Unlock()

	if q.tail == nil {
		q.tail = node
		q.head = node
	} else {
		q.tail.next = node
		q.tail = node
	}
}

// Pop 移除
func (q *NonBlockQueue) Pop() (int, error) {
	// TODO: using cas instead of lock
	q.popLock.Lock()
	defer q.popLock.Unlock()

	if q.head == nil {
		return 0, errors.New("empty queue")
	}
	node := q.head
	q.head = q.head.next

	return node.Val, nil
}

// Dump 打印全部
func (q *NonBlockQueue) Dump() (res []int) {
	it := q.head
	for it != nil {
		res = append(res, it.Val)
		it = it.next
	}
	return
}
