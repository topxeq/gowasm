package main

//go:wasmexport add
func add(x, y uint32) uint32 {
	return x + y
}

// main is required for the `wasi` target, even if it isn't used.
// See https://gowasm.io/languages/tinygo/#why-do-i-have-to-define-main
func main() {}
