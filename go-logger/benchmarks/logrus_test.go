// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package logger provides a simple, lightweight logging library for Go.
package benchmarks

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func Benchmark_Logrus_Info_Time(b *testing.B) {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	for n := 0; n < b.N; n++ {
		log.Info(getMessage(n))
	}
}

func Benchmark_Logrus_Info_Alloc(b *testing.B) {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	var totalAllocs int64
	for n := 0; n < b.N; n++ {
		b.ReportAllocs()
		totalAllocs += int64(testing.AllocsPerRun(1, func() {
			log.Info(getMessage(n))
		}))
	}
	b.SetBytes(totalAllocs / int64(b.N))
}

func Benchmark_Logrus_format_Time(b *testing.B) {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	for n := 0; n < b.N; n++ {
		log.Infof("1:(%d)-2:(%d)-3:(%s)-4:(%s)-5:(%v)-6:(%v)-7:(%v)-8:(%v)-9:(%v)-10:(%v)", fakeFmtArgs()...)
	}
}

func Benchmark_Logrus_format_Alloc(b *testing.B) {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	var totalAllocs int64
	for n := 0; n < b.N; n++ {
		b.ReportAllocs()
		totalAllocs += int64(testing.AllocsPerRun(1, func() {
			log.Infof("1:(%d)-2:(%d)-3:(%s)-4:(%s)-5:(%v)-6:(%v)-7:(%v)-8:(%v)-9:(%v)-10:(%v)", fakeFmtArgs()...)
		}))
	}
	b.SetBytes(totalAllocs / int64(b.N))
}
