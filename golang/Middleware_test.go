package golang

import (
	"github.com/aaronzjc/leetcode/tools"
	"testing"
)

func TestMiddleware(t *testing.T) {
	ctx := new(Context)
	ctx.init()

	var input []string
	ctx.handlers = append(ctx.handlers, func(ctx *Context) {
		input = append(input, "bf-a")
		ctx.next()
		input = append(input, "af-a")
	}, func(ctx *Context) {
		input = append(input, "bf-b")
		ctx.next()
		input = append(input, "af-b")
	}, func(ctx *Context) {
		ctx.next()
		input = append(input, "af-c1")
		input = append(input, "af-c2")
	})
	ctx.next()

	expect := []string{"bf-a", "bf-b", "af-c1", "af-c2", "af-b", "af-a"}
	if !tools.IsArrEquals(input, expect, true) {
		t.Log(input, expect)
		t.Fatal("failed !")
	}

	t.Log("Go Middleware passed !")
}
