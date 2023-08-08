// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package logger provides a simple, lightweight logging library for Go.
package main

import (
	"github.com/pokeyaro/gopkg/go-logger"
)

type User struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Roles []string `json:"roles"`
}

var (
	_str    = "hello~"
	_int    = 123
	_bool   = true
	_bytes  = []byte{'h', 'i'}
	_map    = map[string]string{"aaa": "123", "bbb": "456"}
	_struct = User{Name: "John", Age: 30, Roles: []string{"Admin", "User"}}
)

func egWithAnyType() {
	// You can choose dev or prod to start
	log := logger.SetupDev()

	// Support printing multiple data types
	log.Debug(_str)
	log.Debug(_int)
	log.Debug(_bool)
	log.Debug(_bytes)
	log.Debug(_map)
	log.Debug(&_struct)
}

func egWithPrintfStyle() {
	// You can choose dev or prod to start
	log := logger.SetupDev()

	// Support content format
	log.Tracef("trace %s %d", "level: #", 1)
	log.Debugf("debug %s %d", "level: #", 2)
	log.Infof("info %s %d", "level: #", 3)
	log.Noticef("notice %s %d", "level: #", 4)
	log.Warnf("warn %s %d", "level: #", 5)
	log.Errorf("error %s %d", "level: #", 6)
	// log.Fatalf("fatal %s %d", "level: #", 6)
}

func egWithJson() {
	// The Json method can only be created using the New() function
	log := logger.New().SetRecordToFile(
		&logger.FileRecord{ShouldRec: true, FilePath: "/tmp/logs/", Trigger: logger.LevelWarn},
	)

	// Support printing json
	log.Json(logger.LevelWarn, _struct)
}

func main() {
	egWithAnyType()

	egWithPrintfStyle()

	egWithJson()
}
