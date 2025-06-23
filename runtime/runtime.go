package runtime

import (
	"os"
	"rogchap.com/v8go"
)

var sharedIsolate *v8go.Isolate

func init() {
	sharedIsolate = v8go.NewIsolate() // create once
}

// Runtime holds the context pool
type Runtime struct {
	pool *ContextPool
}

// NewRuntime creates a JS runtime
func NewRuntime() (*Runtime, error) {
	global := v8go.NewObjectTemplate(sharedIsolate)
	InjectConsole(global, sharedIsolate)

	pool := NewContextPool(sharedIsolate, global)
	return &Runtime{pool: pool}, nil
}

func (rt *Runtime) RunScript(code string, filename string) (string, error) {
	ctx := rt.pool.Get()
	defer rt.pool.Put(ctx)

	val, err := ctx.RunScript(code, filename)
	if err != nil {
		return "", err
	}
	return val.String(), nil
}

func (rt *Runtime) RunFile(path string) (string, error) {
	code, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return rt.RunScript(string(code), path)
}
