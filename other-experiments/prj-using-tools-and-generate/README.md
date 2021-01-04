# prj-using-tools-and-generate

An experiment project that uses [tool dependency tracking][gomod-tools] and
`go generate`.

See also:

- [go mod tools example][gomod-tools-example]

## Usage

Compare:

``` sh
❯ rm worldstatus_string.go 
❯ go run .
Hello 3 World
```

and then:

``` sh
❯ go generate
❯ go run .
Hello Beautiful World
```

so `go generate` has to be run manually, but tool dependencies are tracked in
`go.mod` and should be installed automatically.

Note: if it doesn't install the binary to `$GOBIN` automatically, there's a
trick - you can use `go run $MODULE` in the `//go:generate` directive instead
of binary and then you don't depend on the binary being in PATH. See `main.go`.

[gomod-tools]: https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
[gomod-tools-example]: https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md
