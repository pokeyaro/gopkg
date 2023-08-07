# Go-Logger

[English](README.md) | 中文

一个简单封装的轻量级日志库


## 安装

```bash
go get -u github.com/pokeyaro/gopkg/logger
```


## 依赖

- Go 1.18+
- Linux / MacOS / Windows（需要 Go1.18 以上）


## 特色

- 简洁易用的 `API`
- 支持 `JSON` 格式化输出
- 日志级别着色显示


## 快速开始

1. 提供了结构化 `API` 和 `printf` 风格的 `API`

```go
package main

import (
	"github.com/pokeyaro/gopkg/logger"
)

func main() {
	// 还支持 SetupDev 和 SetupProd 模式
	log := logger.New()
	
	// Normal 风格
	log.Debug("hello~")
	
	// Printf-Style 风格
	log.Infof("info %s %d %v", "msg", 1, true)
	
	// Json-Marshal 风格
	log.Json(logger.LevelWarn, User{Name: "John", Age: 30, Roles: []string{"Admin", "User"}})
}
```

2. `go-logger` 允许在以下级别进行日志记录（从最高到最低）：

```textmate
LevelFatal  #5  ColorRed  /  exit(1)
LevelError  #4  ColorMagenta
LevelWarn   #3  ColorYellow
LevelInfo   #2  ColorBlue
LevelDebug  #1  ColorRed
```


## 基准测试

> 在下面表格中，我们列出了几大流行日志库的性能比较结果。时间以纳秒（ns）为单位表示，对象分配数表示每次操作时分配的内存对象数量。

记录一行日志：

| 库         | 时间 (ns/op) | 分配的对象数        |
|-----------|--------------|-------------------|
| zap       | 538 ns/op    | 4 allocs/op       |
| zerolog   | 5167 ns/op   | 0 allocs/op       |
| logrus    | 5048 ns/op   | 24 allocs/op      |
| go-logger | 5098 ns/op   | 32 allocs/op      |

`printf` 风格 `10` 个字段：

| 库         | 时间 (ns/op)  | 分配的对象数       |
|-----------|---------------|------------------|
| zap       | 6601 ns/op    | 116 allocs/op    |
| zerolog   | 21526 ns/op   | 116 allocs/op    |
| logrus    | 22573 ns/op   | 140 allocs/op    |
| go-logger | 21484 ns/op   | 142 allocs/op    |
