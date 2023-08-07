// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

package logger

func ExampleEntry_Debug() {
	log := SetupProd()

	log.Debug("This is the Debug method")

	// Output:
}

func ExampleEntry_Debugf() {
	log := SetupProd()

	log.Debug("This is the %s method", "Debugf")

	// Output:
}

func ExampleEntry_Info() {
	log := SetupProd()

	log.Info("This is the Info method")

	// Output:
}

func ExampleEntry_Infof() {
	log := SetupProd()

	log.Infof("This is the %s method", "Infof")

	// Output:
}

func ExampleEntry_Warn() {
	log := SetupProd()

	log.Warn("This is the Warn method")

	// Output:
}

func ExampleEntry_Warnf() {
	log := SetupProd()

	log.Warnf("This is the %s method", "Warnf")

	// Output:
}

func ExampleEntry_Error() {
	log := SetupProd()

	log.Error("This is the Error method")

	// Output:
}

func ExampleEntry_Errorf() {
	log := SetupProd()

	log.Errorf("This is the %s method", "Errorf")

	// Output:
}

func ExampleEntry_Fatal() {
	log := SetupProd()

	log.Fatal("This is the Fatal method")

	// Output:
}

func ExampleEntry_Fatalf() {
	log := SetupProd()

	log.Fatalf("This is the %s method", "Fatalf")

	// Output:
}

func ExampleEntry_Json() {
	log := SetupProd()

	log.Json(LevelInfo, map[string]any{"name": "John", "age": 30, "roles": []string{"Developer", "DBA"}})

	// Output:
}
