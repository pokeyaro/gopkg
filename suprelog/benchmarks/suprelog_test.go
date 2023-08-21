// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package benchmarks

import (
	"testing"

	"github.com/pokeyaro/gopkg/suprelog"
)

func BenchmarkEntryInfo_Time(b *testing.B) {
	slog := suprelog.New()
	for n := 0; n < b.N; n++ {
		// Code logic to be tested.
		slog.Info(getMessage(n))
	}
}

func BenchmarkEntryInfo_Alloc(b *testing.B) {
	b.ResetTimer()
	slog := suprelog.New()
	var totalAllocs int64
	for n := 0; n < b.N; n++ {
		// Code logic to be tested.
		b.ReportAllocs()
		totalAllocs += int64(testing.AllocsPerRun(1, func() {
			// Perform operations for each iteration here.
			slog.Info(getMessage(n))
		}))
	}
	b.SetBytes(totalAllocs / int64(b.N))
}

func BenchmarkEntryInfof_Time(b *testing.B) {
	slog := suprelog.New()
	for n := 0; n < b.N; n++ {
		// Code logic to be tested.
		slog.Infof("1:(%d)-2:(%d)-3:(%s)-4:(%s)-5:(%v)-6:(%v)-7:(%v)-8:(%v)-9:(%v)-10:(%v)", fakeFmtArgs()...)
	}
}

func BenchmarkEntryInfof_Alloc(b *testing.B) {
	b.ResetTimer()
	slog := suprelog.New()

	var totalAllocs int64
	for n := 0; n < b.N; n++ {
		// Code logic to be tested.
		b.ReportAllocs()
		totalAllocs += int64(testing.AllocsPerRun(1, func() {
			// Perform operations for each iteration here.
			slog.Infof("1:(%d)-2:(%d)-3:(%s)-4:(%s)-5:(%v)-6:(%v)-7:(%v)-8:(%v)-9:(%v)-10:(%v)", fakeFmtArgs()...)
		}))
	}
	b.SetBytes(totalAllocs / int64(b.N))
}
