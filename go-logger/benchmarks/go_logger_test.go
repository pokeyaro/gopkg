// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

package benchmarks

import (
	"testing"

	"github.com/pokeyaro/gopkg/go-logger"
)

func BenchmarkEntry_Info_Time(b *testing.B) {
	log := logger.New()
	for n := 0; n < b.N; n++ {
		// Code logic to be tested.
		log.Info(getMessage(n))
	}
}

func BenchmarkEntry_Info_Alloc(b *testing.B) {
	b.ResetTimer()
	log := logger.New()
	var totalAllocs int64
	for n := 0; n < b.N; n++ {
		// Code logic to be tested.
		b.ReportAllocs()
		totalAllocs += int64(testing.AllocsPerRun(1, func() {
			// Perform operations for each iteration here.
			log.Info(getMessage(n))
		}))
	}
	b.SetBytes(totalAllocs / int64(b.N))
}

func BenchmarkEntry_Infof_Time(b *testing.B) {
	log := logger.New()
	for n := 0; n < b.N; n++ {
		// Code logic to be tested.
		log.Infof("1:(%d)-2:(%d)-3:(%s)-4:(%s)-5:(%v)-6:(%v)-7:(%v)-8:(%v)-9:(%v)-10:(%v)", fakeFmtArgs()...)
	}
}

func BenchmarkEntry_Infof_Alloc(b *testing.B) {
	b.ResetTimer()
	log := logger.New()
	var totalAllocs int64
	for n := 0; n < b.N; n++ {
		// Code logic to be tested.
		b.ReportAllocs()
		totalAllocs += int64(testing.AllocsPerRun(1, func() {
			// Perform operations for each iteration here.
			log.Infof("1:(%d)-2:(%d)-3:(%s)-4:(%s)-5:(%v)-6:(%v)-7:(%v)-8:(%v)-9:(%v)-10:(%v)", fakeFmtArgs()...)
		}))
	}
	b.SetBytes(totalAllocs / int64(b.N))
}
