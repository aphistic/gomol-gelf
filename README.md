gomol-gelf
============

[![GoDoc](https://godoc.org/github.com/aphistic/gomol-gelf?status.svg)](https://godoc.org/github.com/aphistic/gomol-gelf)
[![Build Status](https://img.shields.io/travis/aphistic/gomol-gelf.svg)](https://travis-ci.org/aphistic/gomol-gelf)
[![Code Coverage](https://img.shields.io/codecov/c/github/aphistic/gomol-gelf.svg)](http://codecov.io/github/aphistic/gomol-gelf?branch=master)

gomol-gelf is a logger for [gomol](https://github.com/aphistic/gomol) to support logging to GELF endpoints.

Installation
============

The recommended way to install is via http://gopkg.in

    go get gopkg.in/aphistic/gomol-gelf.v0
    ...
    import "gopkg.in/aphistic/gomol-gelf.v0"

gomol-gelf can also be installed the standard way as well

    go get github.com/aphistic/gomol-gelf
    ...
    import "github.com/aphistic/gomol-gelf"

Examples
========

For brevity a lot of error checking has been omitted, be sure you do your checks!

This is a super basic example of adding a GELF logger to gomol and then logging a few messages:

```go
package main

import (
	"github.com/aphistic/gomol"
	gg "github.com/aphistic/gomol-gelf"
)

func main() {
	// Add an io.Writer logger
	gelfCfg := gg.NewGelfLoggerConfig()
	gelfLogger, _ := gg.NewGelfLogger(gelfCfg)
	gomol.AddLogger(gelfLogger)

	// Set some global attrs that will be added to all
	// messages automatically
	gomol.SetAttr("facility", "gomol.example")
	gomol.SetAttr("another_attr", 1234)

	// Initialize the loggers
	gomol.InitLoggers()
	defer gomol.ShutdownLoggers()

	// Log some debug messages with message-level attrs
	// that will be sent only with that message
	for idx := 1; idx <= 10; idx++ {
		gomol.Dbgm(
			gomol.NewAttrs().
				SetAttr("msg_attr1", 4321),
			"Test message %v", idx)
	}
}

```
