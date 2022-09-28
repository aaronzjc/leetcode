package golang

// 一种中间件实现

// Handler 处理方法
type Handler func(ctx *Context)

// Context 上下文
type Context struct {
	handlers []Handler
	idx      int
}

// init 初始化
func (ctx *Context) init() {
	ctx.idx = -1
}

// next 下一个处理
func (ctx *Context) next() {
	ctx.idx++
	for ; ctx.idx < len(ctx.handlers); ctx.idx++ {
		ctx.handlers[ctx.idx](ctx)
	}
}
