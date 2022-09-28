package golang

import (
	"strconv"
)

// 用协程的方式交替打印数字和字母

// TurnPrint 交替打印
func TurnPrint(counter int) string {
	chNum := make(chan struct{})
	chChar := make(chan struct{})
	done := make(chan struct{})

	res := ""

	// 打印数字
	go func() {
		i := 1
		for {
			select {
			case <-chNum:
				if i%10 == 0 {
					i = 1
				}
				res += strconv.Itoa(i)
				i++
				chChar <- struct{}{}
			}
		}
	}()

	// 打印字母
	go func() {
		i := int('A')
		for {
			select {
			case <-chChar:
				if i > int('Z') {
					i = int('A')
				}
				res += string(rune(i))
				if len(res) >= 2*counter {
					done <- struct{}{}
				} else {
					i++
					chNum <- struct{}{}
				}
			}
		}
	}()

	// 先打印数字
	chNum <- struct{}{}
	<-done

	return res
}
