package golang

import (
	"strconv"
	"testing"
)

func TestConsisHash(t *testing.T) {
	hs := NewConsisHash(3, func(key []byte) uint32 {
		i, err := strconv.Atoi(string(key))
		if err != nil {
			panic(err)
		}
		return uint32(i)
	})

	// 测试空
	if hs.Get("1") != "" {
		t.Log("expect empty")
		t.Fatal("failed !")
	}

	hs.Add("2", "4", "6", "8")
	seeds := map[string]string{
		"0":  "2",
		"20": "2",
		"22": "4",
		"80": "8",
		"90": "2",
	}

	for k, v := range seeds {
		expect := hs.Get(k)
		if v != expect {
			t.Log(expect, v)
			t.Fatal("failed !")
		}
	}

	t.Log("ConsisHash passed !")
}

func TestConsisHash2(t *testing.T) {
	hash1 := NewConsisHash(1, nil)
	hash2 := NewConsisHash(1, nil)

	hash1.Add("Bill", "Bob", "Bonny")
	hash2.Add("Bob", "Bonny", "Bill")

	if hash1.Get("Ben") != hash2.Get("Ben") {
		t.Error("Ben not equals")
	}
}
