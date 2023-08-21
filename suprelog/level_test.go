// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package suprelog

import (
	"fmt"
	"testing"
)

func TestLevel_String(t *testing.T) {
	fmt.Println(LevelTrace)
	fmt.Println(LevelDebug)
	fmt.Println(LevelInfo)
	fmt.Println(LevelNotice)
	fmt.Println(LevelWarn)
	fmt.Println(LevelError)
	fmt.Println(LevelFatal)
}

func TestLevel_Level(t *testing.T) {
	fmt.Println(LevelTrace.Level()) // DEBUG-4
	fmt.Println(LevelDebug.Level())
	fmt.Println(LevelInfo.Level())
	fmt.Println(LevelNotice.Level()) // INFO+2
	fmt.Println(LevelWarn.Level())
	fmt.Println(LevelError.Level())
	fmt.Println(LevelFatal.Level()) // ERROR+4
}
