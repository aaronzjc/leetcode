package hard

import "testing"

func TestJumpGameTwo(t *testing.T)  {
	seeds := []struct{
		input []int
		expect int
	} {
		{[]int{2,3,1,4}, 2},
	}

	for _, v := range seeds {
		result := jump(v.input)
		if result != v.expect {
			t.Errorf(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("JumpTwo passed !")
}
