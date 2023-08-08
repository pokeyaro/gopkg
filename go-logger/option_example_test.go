// Copyright 2023 Pokeya. All rights reserved.
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

// Package logger provides a simple, lightweight logging library for Go.
package logger

func ExampleEntry_Optional() {
	log := New()

	log.Optional(
		WithLevel(LevelDebug),
		WithTrackAbsPath(true),
		WithTimeFormat(FmtDatetime),
		WithEnableColors(true),
		WithRecordToFile(&FileRecord{
			ShouldRec: true,
			FilePath:  "./tmp/logs/",
			Trigger:   LevelWarn,
		}),
	)

	// Output:
}
