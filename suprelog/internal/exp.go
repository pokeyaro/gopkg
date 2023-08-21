// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package internal

// Ternary simulates a ternary operator by executing the trueCallback function
// if the condition is true, or the falseCallback function if the condition is false.
func Ternary(condition bool, trueCallback, falseCallback func()) {
	if condition {
		trueCallback()
	} else {
		falseCallback()
	}
}
