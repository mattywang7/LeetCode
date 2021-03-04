# LeetCode in Golang

## Create a go module:
```
go mod init example.com/hello
```

The `go mod init` will create a `go.mod` file.
The `go.mod` file only appears in the root of the module.
Packages in subdirectories have import paths consisting of the module path plus to the subdirectory.

For example, if we create a subdirectory `world`,
we would not need to run `go mod init` there.
The package would automatically be recognized as part of the `example.com/hello` module,
with import path `example.com/hello/world`.

## Adding a dependency
The primary motivation for Go modules was to improve the experience of using code written by other developers.

The command `go list -m all` lists the current module and all its dependencies.

`go mod tidy` cleans up unused dependencies.