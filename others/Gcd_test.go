package others

import "testing"

func TestGcd(t *testing.T) {
	seeds := []struct {
		a      int
		b      int
		expect int
	}{
		{1, 3, 1},
		{2, 5, 1},
		{5814, 8874, 306},
	}

	for _, v := range seeds {
		result := gcd(v.a, v.b)
		if v.expect != result {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("Gcd passed !")
}
