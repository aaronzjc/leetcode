package golang

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestNewJobPipeline(t *testing.T) {
	ch := NewJobPipeline(3)
	if cap(ch) != 3 || reflect.ValueOf(ch).Kind() != reflect.Chan {
		t.Errorf("%#v", ch)
		t.Fatal("failed !")
	}

	t.Log("NewJobPipeline passed !")
}

func TestProduceConsume(t *testing.T) {
	tt := struct {
		total         int
		limitProducer int
		limitConsumer int
	}{10, 3, 5}

	ch := NewJobPipeline(3)
	done := make(chan struct{})

	// 开启协程来生产
	producerLimiter := make(chan struct{}, tt.limitProducer)
	for i := 0; i < tt.total; i++ {
		producerLimiter <- struct{}{}
		go func() {
			defer func() {
				<-producerLimiter
			}()
			Produce(ch, ProduceAndConsumeJob{
				data: rand.Int(),
				t:    time.Now(),
			})
		}()
		done <- struct{}{}
	}

	// 开启消费
	cnt := 0
	for i := 0; i < tt.limitConsumer; i++ {
		go func() {
			for {
				Consume(ch, func(job ProduceAndConsumeJob) {
					cnt++
				})
			}
		}()
	}

	<-done
	if cnt != tt.total {
		t.Log(tt, cnt)
		t.Fatal("failed !")
	}
}
