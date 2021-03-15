package golang

// 一种中间件实现

type Handler func(ctx *Context)

type Context struct {
	handlers []Handler
	idx      int
}

func (ctx *Context) init() {
	ctx.idx = -1
}

func (ctx *Context) next() {
	ctx.idx++
	for ; ctx.idx < len(ctx.handlers); ctx.idx++ {
		ctx.handlers[ctx.idx](ctx)
	}
}
