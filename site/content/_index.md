+++
title = "the _zero_ dependency _WebAssembly_ runtime for _Go developers_"
layout = "home"
+++


**WebAssembly** is a way to safely run code compiled in other languages. Runtimes
execute WebAssembly Modules (Wasm), which are most often binaries with a
`.wasm` extension.

**gowasm** is the only zero dependency WebAssembly runtime written in Go.

## Get Started

**Get the gowasm CLI** and run any Wasm binary

```bash
curl https://gowasm.io/install.sh | sh
./bin/gowasm run app.wasm
```

**Embed gowasm** in your Go project and extend any app

```go
import "github.com/tetratelabs/gowasm"

// ...

r := gowasm.NewRuntime(ctx)
defer r.Close(ctx)
mod, _ := r.Instantiate(ctx, wasmAdd)
res, _ := mod.ExportedFunction("add").Call(ctx, 1, 2)
```

-----

## Example

The best way to learn gowasm is by trying one of our [examples][1]. The
most [basic example][2] extends a Go application with an addition function
defined in WebAssembly.

## Why zero?

By avoiding CGO, gowasm avoids prerequisites such as shared libraries or libc,
and lets you keep features like cross compilation. Being pure Go, gowasm adds
only a small amount of size to your binary. Meanwhile, gowasm’s API gives
features you expect in Go, such as safe concurrency and context propagation.

### When can I use this?

You can use gowasm today! gowasm's [1.0 release][3] happened in March 2023, and
is [in use]({{< relref "/community/users.md" >}}) by many projects and
production sites.

You can get the latest version of gowasm like this.
```bash
go get github.com/tetratelabs/gowasm@latest
```

Please give us a [star][4] if you end up using gowasm!

[1]: https://github.com/tetratelabs/gowasm/blob/main/examples
[2]: https://github.com/tetratelabs/gowasm/blob/main/examples/basic
[3]: https://tetrate.io/blog/introducing-gowasm-from-tetrate/
[4]: https://github.com/tetratelabs/gowasm/stargazers
