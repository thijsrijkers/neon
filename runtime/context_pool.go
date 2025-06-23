package runtime

import (
	"sync"
	"rogchap.com/v8go"
)

type ContextPool struct {
	iso   *v8go.Isolate
	pool  sync.Pool
	globs *v8go.ObjectTemplate
}

func NewContextPool(iso *v8go.Isolate, global *v8go.ObjectTemplate) *ContextPool {
	return &ContextPool{
		iso:   iso,
		globs: global,
		pool: sync.Pool{
			New: func() any {
				ctx := v8go.NewContext(iso, global) // returns only *v8go.Context (no error)
				return ctx
			},
		},
	}
}

func (p *ContextPool) Get() *v8go.Context {
	return p.pool.Get().(*v8go.Context)
}

func (p *ContextPool) Put(ctx *v8go.Context) {
	p.pool.Put(ctx)
}
