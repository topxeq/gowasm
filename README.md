# gowasm: the zero dependency WebAssembly runtime for Go developers

[![Go Reference](https://pkg.go.dev/badge/github.com/topxeq/gowasm.svg)](https://pkg.go.dev/github.com/topxeq/gowasm) [![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

WebAssembly is a way to safely run code compiled in other languages. Runtimes
execute WebAssembly Modules (Wasm), which are most often binaries with a `.wasm`
extension.

gowasm is a WebAssembly Core Specification [1.0][1] and [2.0][2] compliant
runtime written in Go. It has *zero dependencies*, and doesn't rely on CGO.
This means you can run applications in other languages and still keep cross
compilation.

Import gowasm and extend your Go application with code written in any language!

## About This Repository

This is a personal fork of [wazero](https://github.com/wazero/wazero), renamed to gowasm for use in my Go projects. The original Apache License 2.0 is preserved.

## Quick Start

### Installation

```bash
go get github.com/topxeq/gowasm@latest
```

### Simple Usage Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/topxeq/gowasm"
)

// Simple WebAssembly module that exports an add function
const simpleWasm = `(module
  (func $add (param i32 i32) (result i32)
    local.get 0
    local.get 1
    i32.add)
  (export "add" (func $add))
`

func main() {
	ctx := context.Background()

	// Create runtime
	r := gowasm.NewRuntime(ctx)
	defer r.Close(ctx)

	// Compile WebAssembly module
	compiled, err := r.CompileModule(ctx, []byte(simpleWasm))
	if err != nil {
		panic(err)
	}
	defer compiled.Close(ctx)

	// Instantiate module
	mod, err := r.InstantiateModule(ctx, compiled, gowasm.NewModuleConfig())
	if err != nil {
		panic(err)
	}
	defer mod.Close(ctx)

	// Get exported function
	add := mod.ExportedFunction("add")

	// Call function
	results, err := add.Call(ctx, 2, 3)
	if err != nil {
		panic(err)
	}

	fmt.Printf("2 + 3 = %d\n", results[0])
}
```

## Example

The best way to learn gowasm is by trying one of our [examples](examples/README.md). The
most [basic example](examples/basic) extends a Go application with an addition
function defined in WebAssembly.

## Runtime

There are two runtime configurations supported in gowasm: _Compiler_ is default:

By default, ex `gowasm.NewRuntime(ctx)`, the Compiler is used if supported. You
can also force the interpreter like so:
```go
r := gowasm.NewRuntimeWithConfig(ctx, gowasm.NewRuntimeConfigInterpreter())
```

### Interpreter
Interpreter is a naive interpreter-based implementation of Wasm virtual
machine. Its implementation doesn't have any platform (GOARCH, GOOS) specific
code, therefore _interpreter_ can be used for any compilation target available
for Go (such as `riscv64`).

### Compiler
Compiler compiles WebAssembly modules into machine code ahead of time (AOT),
during `Runtime.CompileModule`. This means your WebAssembly functions execute
natively at runtime. Compiler is faster than Interpreter, often by order of
magnitude (10x) or more. This is done without host-specific dependencies.

### Conformance

Both runtimes pass WebAssembly Core [1.0][3] and [2.0][4] specification tests
on supported platforms:

|   Runtime   |                 Usage                  | amd64 | arm64 | others |
|:-----------:|:--------------------------------------:|:-----:|:-----:|:------:|
| Interpreter | `gowasm.NewRuntimeConfigInterpreter()` |   ✅   |   ✅   |   ✅    |
|  Compiler   |  `gowasm.NewRuntimeConfigCompiler()`   |   ✅   |   ✅   |   ❌    |

## Support Policy

The below support policy focuses on compatibility concerns of those embedding
gowasm into their Go applications.

### gowasm

gowasm's [1.0 release][8] happened in March 2023, and is [in use][9] by many
projects and production sites.

We offer an API stability promise with semantic versioning. In other words, we
promise to not break any exported function signature without incrementing the
major version. This does not mean no innovation: New features and behaviors
happen with a minor version increment, e.g. 1.0.11 to 1.2.0. We also fix bugs
or change internal details with a patch version, e.g. 1.0.0 to 1.0.1.

You can get the latest version of gowasm like this.
```bash
go get github.com/topxeq/gowasm@latest
```

Please give us a [star][10] if you end up using gowasm!

### Go

gowasm has no dependencies except Go and [`x/sys`][12], so the only source of
conflict in your project's use of gowasm is the Go version.

gowasm follows the same version policy as Go's [Release Policy][5]: two
versions. gowasm will ensure these versions work and bugs are valid if there's
an issue with a current Go version.

### Platform

gowasm has two runtime modes: Interpreter and Compiler. The only supported operating
systems are ones we test, but that doesn't necessarily mean other operating
system versions won't work.

We currently test Linux (Ubuntu and scratch), MacOS and Windows as packaged by
[GitHub Actions][6], as well as nested VMs running on Linux for FreeBSD, NetBSD,
OpenBSD, DragonFly BSD, illumos and Solaris.

We also test cross compilation for many `GOOS` and `GOARCH` combinations.

* Interpreter
  * Linux is tested on amd64 and arm64 (native) as well as riscv64 via emulation.
  * Windows, FreeBSD, NetBSD, OpenBSD, DragonFly BSD, illumos and Solaris are
    tested only on amd64.
  * macOS is tested only on arm64.
* Compiler
  * Linux is tested on amd64 and arm64.
  * Windows, FreeBSD, NetBSD, DragonFly BSD, illumos and Solaris are
    tested only on amd64.
  * macOS is tested only on arm64.

gowasm has no dependencies and doesn't require CGO. This means it can also be
embedded in an application that doesn't use an operating system. This is a main
differentiator between gowasm and alternatives.

We verify zero dependencies by running tests in Docker's [scratch image][7].
This approach ensures compatibility with any parent image.

### macOS code-signing entitlements

If you're developing for macOS and need to code-sign your application,
please read issue [#2393][11].

-----
gowasm is a registered trademark of Tetrate.io, Inc. in the United States and/or other countries

[1]: https://www.w3.org/TR/2019/REC-wasm-core-1-20191205/
[2]: https://www.w3.org/TR/2022/WD-wasm-core-2-20220419/
[3]: https://github.com/WebAssembly/spec/tree/wg-1.0/test/core
[4]: https://github.com/WebAssembly/spec/tree/d39195773112a22b245ffbe864bab6d1182ccb06/test/core
[5]: https://go.dev/doc/devel/release
[6]: https://github.com/actions/virtual-environments
[7]: https://docs.docker.com/develop/develop-images/baseimages/#create-a-simple-parent-image-using-scratch
[8]: https://tetrate.io/blog/introducing-gowasm-from-tetrate/
[9]: https://gowasm.io/community/users/
[10]: https://github.com/gowasm/gowasm/stargazers
[11]: https://github.com/gowasm/gowasm/issues/2393
[12]: https://pkg.go.dev/golang.org/x/sys