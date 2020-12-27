# Errors [![Go Reference](https://pkg.go.dev/badge/github.com/yehan2002/errors.svg)](https://pkg.go.dev/github.com/yehan2002/errors)

A simple package for creating and wrapping errors in golang

## Usage

### Creating an error

```go
const ErrEOF = errors.Error("Unexpected end of file")
```

### Wrapping an error

```go
var err error
var config *os.File
if config, err = os.Open("./config.toml");err != nil{
    return errors.Wrap("error getting config file",err)
}
```
