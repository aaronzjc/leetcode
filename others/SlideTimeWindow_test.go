package others

import (
	"fmt"
	"testing"
	"time"
)

func TestSlideWindow(t *testing.T) {
	sw := NewSlideWindowOfTime(100, 10)
	for i := 0; i < 5; i++ {
		bucket, _ := sw.CurrentBucket()
		bucket.Val++
		fmt.Println(bucket.Start)
		time.Sleep(time.Millisecond * 100)
	}
	// should fill 5 bucket
	if sw.Count() != 5 {
		t.Fatal("failed")
	}
	time.Sleep(time.Second)
	// should clean all bucket
	if sw.Count() != 0 {
		t.Fatal("failed")
	}
	// should replace old bucket
	bucket, _ := sw.CurrentBucket()
	bucket.Val++
	if sw.Count() != 1 {
		t.Fatal("failed")
	}

	// 测试Bucket时间异常的问题。理论上在并发情况会发生，但是不好模拟并发。
	sw = NewSlideWindowOfTime(100, 1)
	sw.Buckets[0].Start = time.Now().UnixMilli()
	if _, err := sw.CurrentBucket(); err == nil {
		t.Fatal("failed")
	}

	t.Log("SlideTimeWindow passed")
}
