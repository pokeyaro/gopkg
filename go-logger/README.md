# Go-Logger

English | [中文](README_ZH_CN.md)

A simple, lightweight logging library.


## Installation

```bash
go get -u github.com/pokeyaro/gopkg/go-logger@master
```


## Dependencies

- Go 1.18+
- Linux / MacOS / Windows (requires Go 1.18+)


## Features

- Simple and easy-to-use API
- Supports JSON formatted output
- Log level color highlighting


## Quick Start

1. Provides structured API and printf-style API.

```go
package main

import (
	"github.com/pokeyaro/gopkg/go-logger"
)

func main() {
	// Supports SetupDev and SetupProd modes
	log := logger.New()

	// Normal style
	log.Debug("hello~")

	// Printf-style
	log.Infof("info %s %d %v", "msg", 1, true)

	// Json-Marshal style
	log.Json(logger.LevelWarn, User{Name: "John", Age: 30, Roles: []string{"Admin", "User"}})
}
```

2. `go-logger` allows logging at the following levels (from highest to lowest):

| NO. | Level | Color   | Remark   |
|-----|-------|---------|----------|
| #7  | Fatal | Magenta | exit(1)  |
| #6  | Error | Red     |          |
| #5  | Warn | Yellow  |          |
| #4  | Notice | Green   |          |
| #3  | Info | Cyan    |          |
| #2  | Debug |  Blue   |          |
| #1  | Trace |  Blue   |          |


## Use Case

[example](./example/example.go)


## Benchmarking

> In the table below, we provide performance comparison results for several popular logging libraries. Time is measured in nanoseconds (ns), and object allocations represent the number of memory objects allocated per operation.

Logging a single line:

| Library   | Time (ns/op) | Object Allocations |
|-----------|--------------|--------------------|
| zap       | 538 ns/op    | 4 allocs/op        |
| zerolog   | 5167 ns/op   | 0 allocs/op        |
| logrus    | 5048 ns/op   | 24 allocs/op       |
| go-logger | 5098 ns/op   | 32 allocs/op       |

Printf-style with 10 fields:

| Library | Time (ns/op) | Object Allocations |
|---------|--------------|--------------------|
| zap     | 6601 ns/op   | 116 allocs/op      |
| zerolog | 21526 ns/op  | 116 allocs/op      |
| logrus  | 22573 ns/op  | 140 allocs/op      |
| go-logger  | 21484 ns/op  | 142 allocs/op      |
