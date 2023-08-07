// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

package benchmarks

import (
	"fmt"
	"testing"

	"go.uber.org/zap"
)

func Benchmark_Zap_Info_Time(b *testing.B) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	for n := 0; n < b.N; n++ {
		logger.Info(getMessage(n))
	}
}

func Benchmark_Zap_Info_Alloc(b *testing.B) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	var totalAllocs int64
	for n := 0; n < b.N; n++ {
		b.ReportAllocs()
		totalAllocs += int64(testing.AllocsPerRun(1, func() {
			logger.Info(getMessage(n))
		}))
	}
	b.SetBytes(totalAllocs / int64(b.N))
}

func Benchmark_Zap_format_Time(b *testing.B) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	for n := 0; n < b.N; n++ {
		logger.Info(fmt.Sprintf("1:(%d)-2:(%d)-3:(%s)-4:(%s)-5:(%v)-6:(%v)-7:(%v)-8:(%v)-9:(%v)-10:(%v)", fakeFmtArgs()...))
	}
}

func Benchmark_Zap_format_Alloc(b *testing.B) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	var totalAllocs int64
	for n := 0; n < b.N; n++ {
		b.ReportAllocs()
		totalAllocs += int64(testing.AllocsPerRun(1, func() {
			logger.Info(fmt.Sprintf("1:(%d)-2:(%d)-3:(%s)-4:(%s)-5:(%v)-6:(%v)-7:(%v)-8:(%v)-9:(%v)-10:(%v)", fakeFmtArgs()...))
		}))
	}
	b.SetBytes(totalAllocs / int64(b.N))
}
