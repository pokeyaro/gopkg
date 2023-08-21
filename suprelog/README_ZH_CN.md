<div align=center>
  <img src="assets/logo.png" width="600" height="275" alt="suprelog" /><br/>

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.21-9cf)
![Release](https://img.shields.io/badge/release-1.0.0-green.svg)
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/pokeyaro/gopkg/suprelog)
[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/pokeyaro/gopkg/master/LICENSE)
</div>

[English](README.md) | 中文


## 由来

`"Suprelog"` 这个词汇是由 `"supreme"` 衍生而来，我们希望这个库能够带给您无上至尊的体验，同时还提供了极高的可定制性！

随着 `Go 1.21` 版本中 `slog` 加入标准库，我们基于其后端 `Handler` 标准接口进行了重新实现。
这使得我们能够构建一个功能丰富、灵活可扩展的日志库。通过 `Suprelog`，您可以轻松地管理和记录日志信息，以满足不同场景的需求。
无论是日常开发、调试还是生产环境中，`Suprelog` 都能够为您提供强大的日志处理能力。


## 安装

```bash
go get -u github.com/pokeyaro/gopkg/suprelog@master
```


## 依赖

- `Go 1.21+`
- `Linux` / `MacOS` / `Windows`（需要 `Go1.21` 以上）


## 特性

- 🎭 提供了两套用户友好的 `API`：`Logger-style` 和 `Classical-style`
- 🪶 能够输出文本 `Text` 或 `JSON` 格式的数据
- 🎨 支持自定义 `Level` 级别的背景色阶方案，可选择 `RGB` 或 `HEX` 方式渲染
- 🧩 内置多样的 `Handler`，支持通过 `Chainable Setter` 或 `Option` 加载用户配置
- 🎱 提供了针对 `Fatal` 级别日志的钩子回调方法


## 快速开始

### 最简单的日志示例

```go
package main

import (
    "github.com/pokeyaro/gopkg/suprelog"
)

func main() {
    log := suprelog.New()
    log.Info("hello world")

    // Output:
    // [2023-08-21] | hello world
}
```

### 使用默认初始化日志

```go
package main

import (
    "github.com/pokeyaro/gopkg/suprelog"
)

func main() {
    log := suprelog.Default()
    log.Info("hello world")

    // Output:
    // [2023-08-21 00:03:59.857] [INFO] suprelog/example/main.go:9 | "msg":"hello world"
}
```

### 使用不同的日志级别

```go
package main

import (
    "github.com/pokeyaro/gopkg/suprelog"
)

func main() {
    log := suprelog.Default()
    log.Trace("trace level")
    log.Debug("debug level")
    log.Info("info level")
    log.Notice("notice level")
    log.Warn("warn level")
    log.Error("error level")
    log.Fatal("fatal level")

    // Output:
    // [2023-08-21 00:03:59.857] [DEBUG] suprelog/example/main.go:10 | "msg":"debug level"
    // [2023-08-21 00:03:59.857] [INFO] suprelog/example/main.go:11 | "msg":"info level"
    // [2023-08-21 00:03:59.857] [NOTICE] suprelog/example/main.go:12 | "msg":"notice level"
    // [2023-08-21 00:03:59.857] [WARN] suprelog/example/main.go:13 | "msg":"warn level"
    // [2023-08-21 00:03:59.857] [ERROR] suprelog/example/main.go:14 | "msg":"error level"
    // [2023-08-21 00:03:59.857] [FATAL] suprelog/example/main.go:15 | "msg":"fatal level"
    // Process finished with the exit code 1
}
```

提示：默认为 `Debug` 级别，`Trace` 级别默认不打印。

### Printf 风格与携带 Ctx 上下文方式

```go
package main

import (
    "context"

    "github.com/pokeyaro/gopkg/suprelog"
)

func main() {
    log := suprelog.Default()
    log.Info("hello world")
    log.Infof("hello %s", "world")
    log.InfoCtx(context.TODO(), "hello world")

    // Output:
    // [2023-08-21 00:03:59.857] [INFO] suprelog/example/main.go:11 | "msg":"hello world"
    // [2023-08-21 00:03:59.857] [INFO] suprelog/example/main.go:12 | "msg":"hello world"
    // [2023-08-21 00:03:59.857] [INFO] suprelog/example/main.go:13 | "msg":"hello world"
}
```

### KV 结构化日志 text 与 json 格式

```go
package main

import (
    "github.com/pokeyaro/gopkg/suprelog"
)

func main() {
    tlog := suprelog.ConsoleHandler().InitLogger()
    tlog.Info("hello world", "name", "John", "age", 30, "is_male", true)

    jlog := suprelog.ConsoleHandler().ToggleTypMode().InitLogger()
    jlog.Info("hello world", "name", "John", "age", 30, "is_male", true)

    // Output:
    // [2023-08-21 00:03:59.857] [INFO] suprelog/example/main.go:9 | "msg":"hello world" | "text":"name=John age=30 is_male=true"
    // [2023-08-21 00:03:59.857] [INFO] suprelog/example/main.go:12 | "msg":"hello world" | "json":{"name":"John","age":"30","is_male":"true"}
}
```

### 可定制的 HandlerOptions

```go
package main

import (
    "context"
    "fmt"
    "log/slog"
    "os"
    "time"

    "github.com/pokeyaro/gopkg/suprelog"
)

func main() {
    logger := suprelog.HandlerOptions(
        suprelog.WithWriter(os.Stdout),
        suprelog.WithBuiltinSort([]string{suprelog.FieldLevel, suprelog.FieldTime, suprelog.FieldPos}),
        suprelog.WithExitCode(99),
        suprelog.WithAbsPath(false),
        suprelog.WithTimeFormat(time.DateTime),
        suprelog.WithColorful(true),
        suprelog.WithColorScale(suprelog.ColorTheme("arco")),
        suprelog.WithMode(suprelog.NewMode().SetLog(suprelog.ModeSimplify)),
        suprelog.WithFatalHook(func(ctx context.Context, rec slog.Record) error {
            fmt.Println("This is a fatal callback function!")
            return nil
        }),
        suprelog.WithLogLevel(suprelog.LevelTrace),
    ).InitLogger()

    logger.Fatal("hello world")

    // Output:
    // [FATAL] [2023-08-21 00:03:59] suprelog/example/main.go:31 | hello world
    // This is a fatal callback function!
    // Process finished with the exit code 99
}
```

### 可搭配专属你的 Level 色阶方案

```go
package main

import (
    "github.com/pokeyaro/gopkg/suprelog"
)

func main() {
    colorScheme := suprelog.ColorScale{
        IsRGB: true,
        Colors: []suprelog.ColorItem{
            {
                Level: "TRACE",
                RGB:   []int{201, 205, 212},
            },
            {
                Level: "DEBUG",
                RGB:   []int{22, 93, 255},
            },
            {
                Level: "INFO",
                RGB:   []int{20, 201, 201},
            },
            {
                Level: "NOTICE",
                RGB:   []int{0, 180, 42},
            },
            {
                Level: "WARN",
                RGB:   []int{255, 125, 0},
            },
            {
                Level: "ERROR",
                RGB:   []int{245, 63, 63},
            },
            {
                Level: "FATAL",
                RGB:   []int{245, 49, 157},
            },
        },
    }

    log := suprelog.ConsoleHandler().SetColorScale(&colorScheme).ToggleLogColorful().InitLogger()
    log.Info("hello arco")

    // Output:
    // [2023-08-21 00:03:59.857] [INFO] suprelog/example/main.go:43 | "msg":"hello arco"
}
```

提示：详细可参考 [color.go](./color.go) 文件，`ColorTheme` 函数内置了中国色、`arco-design` 版、`ant-design` 版、`element-plus` 版。

### 经典的链式 API

```go
package main

import (
    "context"
    "errors"

    "github.com/pokeyaro/gopkg/suprelog"
)

func main() {
    log := suprelog.ConsoleHandler().InitClassical()

    ctxKey := "key"
    err := errors.New("it's error")
    ctxWithValue := context.WithValue(context.Background(), ctxKey, "value")

    log.Info().
        Str("hello world").
        Str("Hi %s", "~").
        Int(123).
        Obj([]any{"abc", 456, 'o'}).
        Err(err).
        Ctx(ctxWithValue, ctxKey).
        Emit()

    // Output:
    // [2023-08-21 00:03:59.857] [INFO] suprelog/example/main.go:24 | "msg":"hello world - Hi ~ - 123 - [abc 456 111] - it's error - value"
}
```


## 代码示例

[example.go](./example_test.go)

![img](assets/demo.png)


## 更多的 API 函数签名

`Logger/Classical` 初始化

```go
func Default() Logger
func DefaultLogger() Logger
func DefaultClassical() Classical

func New(funcs ...HandlerFunc) Logger
```

`Entry` 实现 `Logger` 接口

```go
func NewEntry(h slog.Handler) *Entry

func (e *Entry) Trace(msg string, args ...any)
func (e *Entry) Tracef(format string, args ...any)
func (e *Entry) TraceCtx(ctx context.Context, msg string, args ...any)
func (e *Entry) Debug(msg string, args ...any)
func (e *Entry) Debugf(format string, args ...any)
func (e *Entry) DebugCtx(ctx context.Context, msg string, args ...any)
func (e *Entry) Info(msg string, args ...any)
func (e *Entry) Infof(format string, args ...any)
func (e *Entry) InfoCtx(ctx context.Context, msg string, args ...any)
func (e *Entry) Notice(msg string, args ...any)
func (e *Entry) Noticef(format string, args ...any)
func (e *Entry) NoticeCtx(ctx context.Context, msg string, args ...any)
func (e *Entry) Warn(msg string, args ...any)
func (e *Entry) Warnf(format string, args ...any)
func (e *Entry) WarnCtx(ctx context.Context, msg string, args ...any)
func (e *Entry) Error(msg string, args ...any)
func (e *Entry) Errorf(format string, args ...any)
func (e *Entry) ErrorCtx(ctx context.Context, msg string, args ...any)
func (e *Entry) Fatal(msg string, args ...any)
func (e *Entry) Fatalf(format string, args ...any)
func (e *Entry) FatalCtx(ctx context.Context, msg string, args ...any)
```

`Classic` 实现 `Classical` 接口

```go
func NewClassic(h slog.Handler) *Classic

func (c *Classic) Level(l Level) *Classic

func (c *Classic) Trace() Classical
func (c *Classic) Debug() Classical
func (c *Classic) Info() Classical
func (c *Classic) Notice() Classical
func (c *Classic) Warn() Classical
func (c *Classic) Error() Classical
func (c *Classic) Fatal() Classical

func (c *Classic) Str(format string, a ...any) Classical
func (c *Classic) Int(i int) Classical
func (c *Classic) Err(err error) Classical
func (c *Classic) Obj(obj any) Classical
func (c *Classic) Ctx(ctx context.Context, contextKey string) Classical
func (c *Classic) Emit()
```

`Handler` 实现 `slog.Handler` 接口

```go
func NewHandler(w io.Writer) *Handler

func (h *Handler) Enabled(_ context.Context, level slog.Level) bool
func (h *Handler) Handle(ctx context.Context, r slog.Record) error
func (h *Handler) WithAttrs(as []slog.Attr) slog.Handler
func (h *Handler) WithGroup(name string) slog.Handler
```

`Handler` 的 `Option` 函数选项

```go
func HandlerOptions(funcs ...HandlerFunc) *Handler

func WithWriter(w io.Writer) HandlerFunc
func WithBuiltinSort(sorts []string) HandlerFunc
func WithLevel(l Level) HandlerFunc
func WithExitCode(code int) HandlerFunc
func WithAbsPath(isAbs bool) HandlerFunc
func WithTimeFormat(timeFmt string) HandlerFunc
func WithColorful(isColorful bool) HandlerFunc
func WithColorScale(colors *ColorScale) HandlerFunc
func WithMode(mode *Mode) HandlerFunc
func WithFatalHook(hook func(ctx context.Context, rec slog.Record) error) HandlerFunc
```

`Handler` 的 `Setter` 方法链

```go
func Dev() *Handler
func Prod() *Handler

func ConsoleHandler() *Handler

func (h *Handler) SetLogLevel(l Level) *Handler
func (h *Handler) SetBuiltinSort(sorts []string) *Handler
func (h *Handler) SetTimeFormat(format string) *Handler
func (h *Handler) SetColorScale(cs *ColorScale) *Handler
func (h *Handler) SetFatalHook(hook func(ctx context.Context, rec slog.Record) error) *Handler

func (h *Handler) ToggleLogPath() *Handler
func (h *Handler) ToggleLogMode() *Handler
func (h *Handler) ToggleTypMode() *Handler
func (h *Handler) ToggleLogColorful() *Handler

func (h *Handler) InitLogger() Logger
func (h *Handler) InitClassical() Classical
```

`Mode` 方法

```go
func NewMode() *Mode
func (m *Mode) SetLog(log int) *Mode
func (m *Mode) SetTyp(typ string) *Mode
```

`Color` 方法

```go
func NewColorScale() *ColorScale
func ColorTheme(theme string) *ColorScale
```

`Level` 等级

```go
func (l Level) String() string
func (l Level) Level() slog.Level
func (l Level) Int() int
```


## 更多的 Entity 值

`Level` 枚举: 用于设置 `Handler.Level`，使用函数 `WithLogLevel, SetLogLevel`

```textmate
LevelTrace
LevelDebug
LevelInfo
LevelNotice
LevelWarn
LevelError
LevelFatal
```

内置字段枚举: 用于设置 `Handler.builtinSort`，使用函数 `WithBuiltinSort, SetBuiltinSort`

```textmate
FieldTime
FieldLevel
FieldPos
```

`Mode` 枚举: 用于设置 `Handler.mode`，使用函数 `SetLog, SetTyp`

```textmate
// 日志输出格式
ModeSimplify
ModeDetail

// 日志输出类型
ModeText
ModeJson
```

`Color` 主题枚举: 用于设置 `Handler.colorScale`，使用函数 `ColorTheme`

```textmate
"arco" | "ant" | "element" | "china"
```


## 基准测试

```textmate
# 使用 New 初始化日志，记录一条消息：
Benchmark_New-10    	                           15337	     75711 ns/op	   0.08 MB/s	    4320 B/op	      12 allocs/op

# 使用 New 初始化日志，记录 printf 样式 10 个字段：
Benchmark_New-10    	                           11457	     90444 ns/op	   0.72 MB/s	    8592 B/op	     130 allocs/op

# 使用 New 初始化日志，记录 kv 样式 10 个字段：
Benchmark_New-10    	                           11643	     94452 ns/op	   0.95 MB/s	   11937 B/op	     180 allocs/op

# 使用 Default 初始化日志，记录一条消息：
Benchmark_Default-10    	                       12025	     94925 ns/op	   0.18 MB/s	    6253 B/op	      34 allocs/op

# 使用 Default 初始化日志，记录 printf 样式 10 个字段：
Benchmark_Default-10    	                       16717	    118595 ns/op	   0.64 MB/s	   13601 B/op	     152 allocs/op

# 使用 Default 初始化日志，记录 kv 样式 10 个字段：
Benchmark_Default-10    	                       22708	    115781 ns/op	   0.88 MB/s	   13905 B/op	     204 allocs/op

# 使用 HandlerOptions 初始化日志，记录一条消息：
Benchmark_HandlerOptions_InitLogger-10      	   14145	     80830 ns/op	   0.24 MB/s	    5405 B/op	      38 allocs/op

# 使用 HandlerOptions 初始化日志，记录 printf 样式 10 个字段：
Benchmark_HandlerOptions_InitLogger-10    	        9685	    108788 ns/op	   0.72 MB/s	   12753 B/op	     156 allocs/op

# 使用 HandlerOptions 初始化日志，记录 kv 样式 10 个字段：
Benchmark_HandlerOptions_InitLogger-10      	   10000	    114179 ns/op	   0.91 MB/s	   13025 B/op	     208 allocs/op

# 使用 DefaultClassical 初始化日志，记录一条消息：
Benchmark_DefaultClassical-10               	   10000	    115148 ns/op	   0.66 MB/s	   13041 B/op	     152 allocs/op

# 使用 DefaultClassical 初始化日志，记录 printf 样式 10 个字段：
Benchmark_DefaultClassical-10    	               10000	    119092 ns/op	   0.64 MB/s	   13041 B/op	     152 allocs/op

# 使用 HandlerOptions + InitClassical 初始化日志，记录 printf 样式 10 个字段：
Benchmark_HandlerOptions_InitClassical-10    	   10000	    120545 ns/op	   0.68 MB/s	   15681 B/op	     164 allocs/op
```

**表现：** 与内置的 `slog` 库相比，`Superlog` 的 `allocs/op` 大约增加了 `13.28%`，这额外的内存开销是因为其增强的功能、可定制性和抽象性，导致了轻微的性能降低。


## 贡献

非常欢迎您的贡献！如果您发现任何改进或需要修复的问题，请随时提交拉取请求。我喜欢包含针对修复或增强的测试用例的拉取请求。

顺便说一下，我很想知道您对 `suprelog` 库的看法。请随时提交问题或给我发送电子邮件；这对我来说非常重要。


## 作者

[Pokeya Boa](https://github.com/pokeyaro)&nbsp;(<a href="mailto:pokeya.mystic@gmail.com">pokeya.mystic@gmail.com</a>)


## 许可证

`Suprelog` 使用 MIT 许可证进行发布，详见 [LICENSE](../LICENSE) 文件。