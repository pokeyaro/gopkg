// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

package main

import (
	"github.com/pokeyaro/gopkg/go-logger"
)

func main() {
	// You can choose dev or prod to start
	log := logger.SetupDev()

	type User struct {
		Name  string   `json:"name"`
		Age   int      `json:"age"`
		Roles []string `json:"roles"`
	}

	// Support printing multiple data types
	log.Debug("hello~")
	log.Debug(123)
	log.Debug(true)
	log.Debug([]rune{'你', '好'})
	log.Debug(map[string]string{"aaa": "123", "bbb": "456"})
	log.Debug(User{Name: "John", Age: 30})

	// Support printing json
	log.Json(logger.LevelWarn, User{Name: "John", Age: 30, Roles: []string{"Admin", "User"}})

	// Support content format
	log.Infof("info %s %d", "msg...", 1)
	log.Warnf("warn %s %d", "msg...", 2)
	log.Errorf("error %s %d", "msg...", 3)
	log.Fatalf("fatal %s...", "end!!!")
}
