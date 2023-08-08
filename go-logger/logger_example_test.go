// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

package logger

import (
	"context"
)

func ExampleEntry_Trace() {
	log := SetupProd()

	log.Trace("This is the Trace method")

	// Output:
}

func ExampleEntry_Tracef() {
	log := SetupProd()

	log.Tracef("This is the %s method", "Tracef")

	// Output:
}

func ExampleEntry_CtxTracef() {
	log := SetupProd()

	ctx := context.TODO()

	log.CtxTracef(ctx, "This is the %s method with context", "Tracef")

	// Output:
}

func ExampleEntry_Debug() {
	log := SetupProd()

	log.Debug("This is the Debug method")

	// Output:
}

func ExampleEntry_Debugf() {
	log := SetupProd()

	log.Debugf("This is the %s method", "Debugf")

	// Output:
}

func ExampleEntry_CtxDebugf() {
	log := SetupProd()

	ctx := context.TODO()

	log.CtxDebugf(ctx, "This is the %s method with context", "Debugf")

	// Output:
}

func ExampleEntry_Info() {
	log := SetupProd()

	log.Info("This is the Info method")

	// Output:
}

func ExampleEntry_Infof() {
	log := SetupProd()

	log.Infof("This is the %s method", "Infof")

	// Output:
}

func ExampleEntry_CtxInfof() {
	log := SetupProd()

	ctx := context.TODO()

	log.CtxInfof(ctx, "This is the %s method with context", "Infof")

	// Output:
}

func ExampleEntry_Notice() {
	log := SetupProd()

	log.Notice("This is the Notice method")

	// Output:
}

func ExampleEntry_Noticef() {
	log := SetupProd()

	log.Noticef("This is the %s method", "Noticef")

	// Output:
}

func ExampleEntry_CtxNoticef() {
	log := SetupProd()

	ctx := context.TODO()

	log.CtxNoticef(ctx, "This is the %s method with context", "Noticef")

	// Output:
}

func ExampleEntry_Warn() {
	log := SetupProd()

	log.Warn("This is the Warn method")

	// Output:
}

func ExampleEntry_Warnf() {
	log := SetupProd()

	log.Warnf("This is the %s method", "Warnf")

	// Output:
}

func ExampleEntry_CtxWarnf() {
	log := SetupProd()

	ctx := context.TODO()

	log.CtxWarnf(ctx, "This is the %s method with context", "Warnf")

	// Output:
}

func ExampleEntry_Error() {
	log := SetupProd()

	log.Error("This is the Error method")

	// Output:
}

func ExampleEntry_Errorf() {
	log := SetupProd()

	log.Errorf("This is the %s method", "Errorf")

	// Output:
}

func ExampleEntry_CtxErrorf() {
	log := SetupProd()

	ctx := context.TODO()

	log.CtxErrorf(ctx, "This is the %s method with context", "Errorf")

	// Output:
}

func ExampleEntry_Fatal() {
	log := SetupProd()

	log.Fatal("This is the Fatal method")

	// Output:
}

func ExampleEntry_Fatalf() {
	log := SetupProd()

	log.Fatalf("This is the %s method", "Fatalf")

	// Output:
}

func ExampleEntry_CtxFatalf() {
	log := SetupProd()

	ctx := context.TODO()

	log.CtxFatalf(ctx, "This is the %s method with context", "Fatalf")

	// Output:
}

func ExampleEntry_Json() {
	log := New()

	log.Json(LevelInfo, map[string]any{"name": "John", "age": 30, "roles": []string{"Developer", "DBA"}})

	// Output:
}
