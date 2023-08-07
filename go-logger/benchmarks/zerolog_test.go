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
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func Benchmark_Zerolog_Info_Time(b *testing.B) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	for n := 0; n < b.N; n++ {
		log.Info().Msg(getMessage(n))
	}
}

func Benchmark_Zerolog_Info_Alloc(b *testing.B) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	var totalAllocs int64
	for n := 0; n < b.N; n++ {
		b.ReportAllocs()
		totalAllocs += int64(testing.AllocsPerRun(1, func() {
			log.Info().Msg(getMessage(n))
		}))
	}
	b.SetBytes(totalAllocs / int64(b.N))
}

func Benchmark_Zerolog_format_Time(b *testing.B) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	for n := 0; n < b.N; n++ {
		log.Info().Msg(fmt.Sprintf("1:(%d)-2:(%d)-3:(%s)-4:(%s)-5:(%v)-6:(%v)-7:(%v)-8:(%v)-9:(%v)-10:(%v)", fakeFmtArgs()...))
	}
}

func Benchmark_Zerolog_format_Alloc(b *testing.B) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	var totalAllocs int64
	for n := 0; n < b.N; n++ {
		b.ReportAllocs()
		totalAllocs += int64(testing.AllocsPerRun(1, func() {
			log.Info().Msg(fmt.Sprintf("1:(%d)-2:(%d)-3:(%s)-4:(%s)-5:(%v)-6:(%v)-7:(%v)-8:(%v)-9:(%v)-10:(%v)", fakeFmtArgs()...))
		}))
	}
	b.SetBytes(totalAllocs / int64(b.N))
}
