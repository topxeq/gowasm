## CLI example

This example shows a simple CLI application compiled to WebAssembly and
executed with the gowasm CLI.

```bash
$ go run github.com/tetratelabs/gowasm/cmd/gowasm run testdata/cli.wasm 3 4
```

The gowasm CLI can run stand-alone Wasm binaries, providing access to any
arguments passed after the path. The Wasm binary reads arguments and otherwise
operates on the host via WASI functions.
