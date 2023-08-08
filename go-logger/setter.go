// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package logger provides a simple, lightweight logging library for Go.
package logger

import (
	"github.com/pokeyaro/gopkg/go-logger/utils"
)

func (entry *Entry) SetLevel(l Level) *Entry {
	entry.level = l
	return entry
}

func (entry *Entry) SetTrackAbsPath(isAbs bool) *Entry {
	entry.trackAbsPath = isAbs
	return entry
}

func (entry *Entry) SetTimeFormat(dt DateFmt) *Entry {
	entry.timeFormat = dt
	return entry
}

func (entry *Entry) SetEnableColors(enable bool) *Entry {
	entry.isColorful = enable
	return entry
}

func (entry *Entry) SetRecordToFile(record RecordRule) *Entry {
	filePath := record.GetPosition()

	if err := utils.Mkdir(filePath); err != nil {
		panic(err.Error())
	}

	entry.recordToFile = record

	return entry
}
