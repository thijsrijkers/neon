package unit

import (
	"testing"

	"rogchap.com/v8go"
)

func TestV8goSimpleEval(t *testing.T) {
	iso := v8go.NewIsolate()
	defer iso.Dispose()

	ctx := v8go.NewContext(iso)

	val, err := ctx.RunScript("1+2", "test.js")
	if err != nil {
		t.Fatalf("Failed to run script: %v", err)
	}

	expected := "3"
	if val.String() != expected {
		t.Errorf("Unexpected result: got %s, want %s", val.String(), expected)
	}
}
