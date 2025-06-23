package runtime

import (
	"fmt"
	"rogchap.com/v8go"
)

func InjectConsole(global *v8go.ObjectTemplate, iso *v8go.Isolate) {
	console := v8go.NewObjectTemplate(iso)

	logFn := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		args := info.Args()          // get all arguments
		for i := 0; i < len(args); i++ {
			fmt.Print(args[i].String(), " ")
		}
		fmt.Println()
		return nil
	})

	console.Set("log", logFn)
	global.Set("console", console)
}
