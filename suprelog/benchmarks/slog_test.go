// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package benchmarks

import (
	"fmt"
	"log/slog"
	"testing"
)

func BenchmarkSlog_Info_Time(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Code logic to be tested.
		slog.Info(getMessage(n))
	}
}

func BenchmarkSlog_Info_Alloc(b *testing.B) {
	b.ResetTimer()
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

func BenchmarkSlog_Infof_Time(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Code logic to be tested.
		slog.Info(fmt.Sprintf("1:(%d)-2:(%d)-3:(%s)-4:(%s)-5:(%v)-6:(%v)-7:(%v)-8:(%v)-9:(%v)-10:(%v)", fakeFmtArgs()...))
	}
}

func BenchmarkSlog_Infof_Alloc(b *testing.B) {
	b.ResetTimer()
	var totalAllocs int64
	for n := 0; n < b.N; n++ {
		// Code logic to be tested.
		b.ReportAllocs()
		totalAllocs += int64(testing.AllocsPerRun(1, func() {
			// Perform operations for each iteration here.
			slog.Info(fmt.Sprintf("1:(%d)-2:(%d)-3:(%s)-4:(%s)-5:(%v)-6:(%v)-7:(%v)-8:(%v)-9:(%v)-10:(%v)", fakeFmtArgs()...))
		}))
	}
	b.SetBytes(totalAllocs / int64(b.N))
}
