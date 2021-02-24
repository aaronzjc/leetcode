package golang

import (
	"math/rand"
	"sync"
	"time"
)

// 使用协程模拟生产者和消费者。具有如下的特点
// 1. 带缓冲的通道
// 2. 限制生产者的并发数

type ProduceAndConsumeJob struct {
	data int
	t    time.Time
}

type Checker struct {
	*sync.Mutex
	sumP   int
	totalP int
	sumC   int
	totalC int
}

func NewJobPipeline(num int) chan ProduceAndConsumeJob {
	return make(chan ProduceAndConsumeJob, num)
}

func ProduceAndConsume(limitC int, limitP int, total int) Checker {
	ch := NewJobPipeline(limitC)
	checker := Checker{
		Mutex: new(sync.Mutex),
	}
	limitProducer := make(chan struct{}, limitP)
	go func() {
		wgp := sync.WaitGroup{}
		for i := 0; i < total; i++ {
			wgp.Add(1)
			limitProducer <- struct{}{}
			go func() {
				defer func() {
					<-limitProducer
					wgp.Done()
				}()
				t := time.Now()
				rand.Seed(t.UnixNano())
				data := rand.Int() % 5
				job := ProduceAndConsumeJob{
					data: data,
					t:    t,
				}

				// 校验
				checker.Lock()
				checker.sumP += job.data
				checker.totalP++
				checker.Unlock()

				ch <- job
			}()
		}
		wgp.Wait()
		close(ch)
		close(limitProducer)
	}()

	wgc := &sync.WaitGroup{}
	wgc.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wgc.Done()
			for {
				job, ok := <-ch
				if !ok {
					break
				}

				// 保证并发安全
				checker.Lock()
				checker.sumC += job.data
				checker.totalC++
				checker.Unlock()
			}
		}()
	}

	wgc.Wait()

	return checker
}
