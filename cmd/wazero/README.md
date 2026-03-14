## gowasm CLI

The gowasm CLI can be used to execute a standalone WebAssembly binary.

### Installation

```bash
$ go install github.com/tetratelabs/gowasm/cmd/gowasm@latest
```

### Usage

The gowasm CLI accepts a single argument, the path to a WebAssembly binary.
Arguments can be passed to the WebAssembly binary itself after the path.

```bash
gowasm run calc.wasm 1 + 2
```

In addition to arguments, the WebAssembly binary has access to stdout, stderr,
and stdin.

