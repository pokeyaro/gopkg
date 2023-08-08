// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

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
