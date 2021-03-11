package golang

import (
	"sync"
	"testing"
)

func TestNoneBlockingQueue(t *testing.T) {
	var inTotal, outTotal, expectTotal int
	c, s := 50, 2000
	expectTotal = c * s
	queue := NewNonBlockQueue()
	wgp := sync.WaitGroup{}
	for i := 0; i < c; i++ {
		wgp.Add(1)
		go func() {
			for i := 0; i < s; i++ {
				queue.Push(i)
			}
			wgp.Done()
		}()
	}
	wgp.Wait()
	inTotal = len(queue.Dump())
	if inTotal != expectTotal {
		t.Log("producer not match", inTotal, expectTotal)
		t.Fatal("failed !")
	}

	lock := &sync.Mutex{}
	wgc := sync.WaitGroup{}
	for i := 1; i < 10; i++ {
		wgc.Add(1)
		go func() {
			for {
				_, err := queue.Pop()
				if err != nil {
					break
				}

				lock.Lock()
				outTotal++
				lock.Unlock()
			}
			wgc.Done()
		}()
	}
	wgc.Wait()

	if inTotal != outTotal || outTotal == 0 {
		t.Log("consumer not match", inTotal, outTotal)
		t.Fatal("failed !")
	}

	_, err := NewNonBlockQueue().Pop()
	if err == nil {
		t.Fatal("failed !")
	}

	t.Log("NonBlockingQueue passed !")
}
