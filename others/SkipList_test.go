package others

import "testing"

func TestSkipList(t *testing.T) {
	seeds := []int{1, 5, 10, 15, 3, 7, 8, 11, 13}
	l := NewSkipList()
	m := make(map[int]*SkipListNode)
	for _, v := range seeds {
		m[v] = l.Add(v)
	}
	l.Dump()
	for _, v := range seeds {
		result, _ := l.Get(v)
		expect := m[v]
		if result.val != expect.val && len(result.level) != len(expect.level) {
			t.Error(result, expect)
			t.Fatal("failed !")
		}
	}

	if _, err := l.Get(100); err == nil {
		t.Error("found 100 error")
	}

	t.Log("SkipList passed !")
}
