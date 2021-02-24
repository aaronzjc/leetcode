package golang

import (
	"reflect"
	"testing"
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
	seeds := []struct {
		lc    int
		lp    int
		total int
	}{
		{3, 3, 100},
		{10, 3, 1000},
		{3, 10, 1000},
		{100, 5, 100000},
	}
	for _, v := range seeds {
		result := ProduceAndConsume(v.lc, v.lp, v.total)
		if result.totalC != result.totalP && result.sumP != result.sumC {
			t.Error(v, result)
			t.Fatal("failed !")
		}
	}

	t.Log("ProduceAndConsume passed !")
}
