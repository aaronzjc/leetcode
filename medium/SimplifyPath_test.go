package medium

import "testing"

func TestSimplifyPath(t *testing.T) {
	seeds := []struct {
		input  string
		expect string
	}{
		{"/home/", "/home"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
		{"/a/./b/../../c/", "/c"},
		{"/a/../../b/../c//.//", "/c"},
		{"/a//b////c/d//././/..", "/a/b/c"},
	}

	for _, v := range seeds {
		result := simplifyPath(v.input)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("SimplifyPath passed !")
}
