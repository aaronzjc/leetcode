package golang

import (
	"time"
)

// 使用协程模拟生产者和消费者

type ProduceAndConsumeJob struct {
	data int
	t    time.Time
}

func NewJobPipeline(num int) chan ProduceAndConsumeJob {
	return make(chan ProduceAndConsumeJob, num)
}

func Produce(pipe chan ProduceAndConsumeJob, job ProduceAndConsumeJob) {
	pipe <- job
}

func Consume(pipe <-chan ProduceAndConsumeJob, handler func(job ProduceAndConsumeJob)) {
	item := <-pipe
	handler(item)
}
