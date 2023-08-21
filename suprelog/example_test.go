// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package suprelog_test

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/pokeyaro/gopkg/suprelog"
)

func ExampleDefault() {
	logger := suprelog.Default()
	logger.Info("hello world")

	// Output:
	// [2023-08-21 00:03:59.857] [INFO] suprelog/example_test.go:10 | "msg":"hello world"
}

func ExampleDefaultLogger() {
	logger := suprelog.DefaultLogger()
	logger.Info("hello world")

	// Output:
	// [2023-08-21 00:03:59.857] [INFO] suprelog/example_test.go:10 | "msg":"hello world"
}

func ExampleDefaultClassical() {
	log := suprelog.DefaultClassical()
	log.Info().Str("hello world").Emit()

	// Output:
	// [2023-08-21 00:03:59.857] [INFO] suprelog/example_test.go:10 | hello world
}

func ExampleNew() {
	logger := suprelog.New()
	logger.Info("hello world")

	// Output:
	// [2023-08-21] | hello world
}

func ExampleHandlerOptions() {
	logger := suprelog.HandlerOptions(
		suprelog.WithWriter(os.Stdout),
		suprelog.WithBuiltinSort([]string{suprelog.FieldTime, suprelog.FieldLevel, suprelog.FieldPos}),
		suprelog.WithExitCode(99),
		suprelog.WithAbsPath(false),
		suprelog.WithTimeFormat(time.DateTime),
		suprelog.WithColorful(true),
		suprelog.WithColorScale(suprelog.ColorTheme("china")),
		suprelog.WithMode(suprelog.NewMode().SetLog(suprelog.ModeSimplify)),
		suprelog.WithFatalHook(func(ctx context.Context, rec slog.Record) error {
			fmt.Println("This is a fatal callback function! You can send alarms, or dump logs to remote ELK, etc.")
			fmt.Printf("Log internal information:\ncontext: %v, record: %v\n", ctx, rec)
			return nil
		}),
		suprelog.WithLogLevel(suprelog.LevelTrace),
	).InitLogger()

	logger.Info("hello world")

	// Output:
	// [2023-08-21 00:03:59.857] [INFO] suprelog/example_test.go:10 | hello world
}

func ExampleEntry_Trace() {
	logger := suprelog.Default()
	logger.Trace("Personnel introduction", "username", "John", "age", 30, "is_male", true)

	// Output:
	//
}

func ExampleEntry_Tracef() {
	logger := suprelog.Default()
	logger.Tracef("Personnel introduction: username %s age %d is_male %v", "John", 30, true)

	// Output:
	//
}

func ExampleEntry_TraceCtx() {
	logger := suprelog.Default()
	logger.TraceCtx(context.TODO(), "Personnel introduction", "username", "John", "age", 30, "is_male", true)

	// Output:
	//
}

func ExampleEntry_Debug() {
	logger := suprelog.Default()
	logger.Debug("Personnel introduction", "username", "John", "age", 30, "is_male", true)

	// Output:
	// [2023-08-21 00:03:59.857] [DEBUG] suprelog/example_test.go:10 | "msg":"Personnel introduction" | "text":"username=John age=30 is_male=true"
}

func ExampleEntry_Debugf() {
	logger := suprelog.Default()
	logger.Debugf("Personnel introduction: username=%s age=%d is_male=%v", "John", 30, true)

	// Output:
	// [2023-08-21 00:03:59.857] [DEBUG] suprelog/example_test.go:10 | "msg":"Personnel introduction: username=John age=30 is_male=true"
}

func ExampleEntry_DebugCtx() {
	logger := suprelog.Default()
	logger.DebugCtx(context.TODO(), "Personnel introduction", "username", "John", "age", 30, "is_male", true)

	// Output:
	// [2023-08-21 00:03:59.857] [DEBUG] suprelog/example_test.go:10 | "msg":"Personnel introduction" | "text":"username=John age=30 is_male=true"
}

func ExampleEntry_Info() {
	logger := suprelog.Default()
	logger.Info("Personnel introduction", "username", "John", "age", 30, "is_male", true)

	// Output:
	// [2023-08-21 00:03:59.857] [INFO] suprelog/example_test.go:10 | "msg":"Personnel introduction" | "text":"username=John age=30 is_male=true"
}

func ExampleEntry_Infof() {
	logger := suprelog.Default()
	logger.Infof("Personnel introduction: username=%s age=%d is_male=%v", "John", 30, true)

	// Output:
	// [2023-08-21 00:03:59.857] [INFO] suprelog/example_test.go:10 | "msg":"Personnel introduction: username=John age=30 is_male=true"
}

func ExampleEntry_InfoCtx() {
	logger := suprelog.Default()
	logger.InfoCtx(context.TODO(), "Personnel introduction", "username", "John", "age", 30, "is_male", true)

	// Output:
	// [2023-08-21 00:03:59.857] [INFO] suprelog/example_test.go:10 | "msg":"Personnel introduction" | "text":"username=John age=30 is_male=true"
}

func ExampleEntry_Notice() {
	logger := suprelog.Default()
	logger.Notice("Personnel introduction", "username", "John", "age", 30, "is_male", true)

	// Output:
	// [2023-08-21 00:03:59.857] [NOTICE] suprelog/example_test.go:10 | "msg":"Personnel introduction" | "text":"username=John age=30 is_male=true"
}

func ExampleEntry_Noticef() {
	logger := suprelog.Default()
	logger.Noticef("Personnel introduction: username=%s age=%d is_male=%v", "John", 30, true)

	// Output:
	// [2023-08-21 00:03:59.857] [NOTICE] suprelog/example_test.go:10 | "msg":"Personnel introduction: username=John age=30 is_male=true"
}

func ExampleEntry_NoticeCtx() {
	logger := suprelog.Default()
	logger.NoticeCtx(context.TODO(), "Personnel introduction", "username", "John", "age", 30, "is_male", true)

	// Output:
	// [2023-08-21 00:03:59.857] [NOTICE] suprelog/example_test.go:159 | "msg":"Personnel introduction" | "text":"username=John age=30 is_male=true"
}

func ExampleEntry_Warn() {
	logger := suprelog.Default()
	logger.Warn("Personnel introduction", "username", "John", "age", 30, "is_male", true)

	// Output:
	// [2023-08-21 00:03:59.857] [WARN] suprelog/example_test.go:10 | "msg":"Personnel introduction" | "text":"username=John age=30 is_male=true"
}

func ExampleEntry_Warnf() {
	logger := suprelog.Default()
	logger.Warnf("Personnel introduction: username=%s age=%d is_male=%v", "John", 30, true)

	// Output:
	// [2023-08-21 00:03:59.857] [WARN] suprelog/example_test.go:10 | "msg":"Personnel introduction: username=John age=30 is_male=true"
}

func ExampleEntry_WarnCtx() {
	logger := suprelog.Default()
	logger.WarnCtx(context.TODO(), "Personnel introduction", "username", "John", "age", 30, "is_male", true)

	// Output:
	// [2023-08-21 00:03:59.857] [WARN] suprelog/example_test.go:10 | "msg":"Personnel introduction" | "text":"username=John age=30 is_male=true"
}

func ExampleEntry_Error() {
	logger := suprelog.Default()
	logger.Error("Personnel introduction", "username", "John", "age", 30, "is_male", true)

	// Output:
	// [2023-08-21 00:03:59.857] [ERROR] suprelog/example_test.go:10 | "msg":"Personnel introduction" | "text":"username=John age=30 is_male=true"
}

func ExampleEntry_Errorf() {
	logger := suprelog.Default()
	logger.Errorf("Personnel introduction: username=%s age=%d is_male=%v", "John", 30, true)

	// Output:
	// [2023-08-21 00:03:59.857] [ERROR] suprelog/example_test.go:10 | "msg":"Personnel introduction: username=John age=30 is_male=true"
}

func ExampleEntry_ErrorCtx() {
	logger := suprelog.Default()
	logger.ErrorCtx(context.TODO(), "Personnel introduction", "username", "John", "age", 30, "is_male", true)

	// Output:
	// [2023-08-21 00:03:59.857] [ERROR] suprelog/example_test.go:10 | "msg":"Personnel introduction" | "text":"username=John age=30 is_male=true"
}

func ExampleEntry_Fatal() {
	logger := suprelog.Default()
	logger.Error("Personnel introduction", "username", "John", "age", 30, "is_male", true)

	// Output:
	// [2023-08-21 00:03:59.857] [FATAL] suprelog/example_test.go:10 | "msg":"Personnel introduction" | "text":"username=John age=30 is_male=true"
}

func ExampleEntry_Fatalf() {
	logger := suprelog.Default()
	logger.Errorf("Personnel introduction: username=%s age=%d is_male=%v", "John", 30, true)

	// Output:
	// [2023-08-21 00:03:59.857] [FATAL] suprelog/example_test.go:10 | "msg":"Personnel introduction: username=John age=30 is_male=true"
}

func ExampleEntry_FatalCtx() {
	logger := suprelog.Default()
	logger.ErrorCtx(context.TODO(), "Personnel introduction", "username", "John", "age", 30, "is_male", true)

	// Output:
	// [2023-08-21 00:03:59.857] [FATAL] suprelog/example_test.go:10 | "msg":"Personnel introduction" | "text":"username=John age=30 is_male=true"
}

func ExampleHandler_InitClassical() {
	logger := suprelog.HandlerOptions().InitClassical()
	logger.Info().Str("hello world").Emit()

	// Output:
	// [2023-08-21] | hello world
}

func ExampleHandler_InitLogger() {
	logger := suprelog.HandlerOptions().InitLogger()
	logger.Info("hello world")

	// Output:
	// [2023-08-21] | hello world
}

func ExampleConsoleHandler() {
	logger := suprelog.ConsoleHandler().
		SetLogLevel(suprelog.LevelTrace).
		SetBuiltinSort([]string{suprelog.FieldLevel, suprelog.FieldPos}).
		ToggleLogColorful().
		ToggleLogMode().
		InitLogger()
	logger.Info("hello world")

	// Output:
	// [INFO] suprelog/example_test.go:10 | hello world
}
