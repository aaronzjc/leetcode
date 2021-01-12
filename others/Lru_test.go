package others

import (
	"testing"
)

func TestLRUCase1(t *testing.T) {
	var dd string
	cache1 := NewLRUCache(3)

	type Eg struct {
		val int
		e   bool
	}
	dumpSeeds := []struct {
		key  int
		val  int
		get  bool
		set  bool
		eg   Eg
		dump string
	}{
		{1, 1, false, true, Eg{}, "1"},             // set   1
		{1, 1, true, false, Eg{1, true}, "1"},      // get   1
		{2, 2, true, false, Eg{0, false}, "1"},     // get   1
		{2, 2, false, true, Eg{}, "2,1"},           // set   2 1
		{3, 3, false, true, Eg{}, "3,2,1"},         // set   3 2 1
		{2, 2, true, false, Eg{2, true}, "2,3,1"},  // get   2 3 1
		{4, 4, false, true, Eg{}, "4,2,3"},         // set   4 2 3
		{5, 5, true, false, Eg{0, false}, "4,2,3"}, // get  4 2 3
		{2, 2, false, true, Eg{}, "2,4,3"},         // set   2 4 3
		{2, 2, false, true, Eg{}, "2,4,3"},         // set   2 4 3
		{3, 3, false, true, Eg{}, "3,2,4"},         // set   3 2 4
		{5, 5, false, true, Eg{}, "5,3,2"},         // set   5 3 2
		{5, 5, true, false, Eg{5, true}, "5,3,2"},  // get  5 3 2
	}

	for _, v := range dumpSeeds {
		if v.get {
			res, ok := cache1.Get(LRUKey(v.key))
			if res != LRUVal(v.eg.val) && ok != v.eg.e {
				t.Fatalf("get key = %v, res = %v, ok = %v, expect = %v", v.key, res, ok, v.eg)
			}
		}
		if v.set {
			cache1.Set(LRUKey(v.key), LRUVal(v.val))
		}
		dd = cache1.Dump()
		if dd != v.dump {
			t.Fatalf("dd = %s, expect = %s", dd, v.dump)
		}
	}

	t.Log("passed !")
}
