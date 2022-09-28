/**
 * 滑动时间窗口
 * 滑动时间窗口用于统计一段时间内，出现的次数等，常用于流量控制场景。
 * 因为时间是流动的，统计的窗口也要跟着流动，所以称为滑动时间窗口。
 */

package others

import (
	"errors"
	"time"
)

type SlideWindowOfTime struct {
	BucketMs     int64
	BucketLength int
	Buckets      []*BucketItem
}
type BucketItem struct {
	Start int64
	Val   int
}

func NewSlideWindowOfTime(bucketMs int64, bucketLength int) *SlideWindowOfTime {
	sw := &SlideWindowOfTime{
		BucketMs:     bucketMs,
		BucketLength: bucketLength,
		Buckets:      make([]*BucketItem, bucketLength),
	}
	now := time.Now().UnixMilli()
	start := now - (now % bucketMs)
	idx := int((now / bucketMs) % int64(bucketLength))
	for i := idx; i < bucketLength; i++ {
		item := &BucketItem{
			Start: start,
			Val:   0,
		}
		sw.Buckets[i] = item
		start = start + bucketMs
	}
	for i := 0; i < idx; i++ {
		item := &BucketItem{
			Start: start,
			Val:   0,
		}
		sw.Buckets[i] = item
		start = start + bucketMs
	}
	return sw
}

func (s *SlideWindowOfTime) CurrentBucket() (*BucketItem, error) {
	now := time.Now().UnixMilli()
	start := now - (now % s.BucketMs)
	idx := int((now / s.BucketMs) % int64(s.BucketLength))
	oldBucket := s.Buckets[idx]
	if oldBucket.Start > start {
		return nil, errors.New("clock error")
	}
	if oldBucket.Start < start {
		// 下标相同的Bucket，可能是经过一轮或者N轮了
		// 理论上其他的Bucket是过期的，但是我们这里不处理。放在统计时进行处理。
		oldBucket.Start = start
		oldBucket.Val = 0
	}
	return oldBucket, nil
}

func (sw *SlideWindowOfTime) Count() int {
	now := time.Now().UnixMilli()
	start := now - (now % sw.BucketMs)
	minStart := start - int64(sw.BucketLength)*sw.BucketMs
	var count int
	for _, v := range sw.Buckets {
		if v.Start < minStart {
			// Bucket是过期的了
			continue
		}
		count += v.Val
	}
	return count
}
