# Logange

[![GoDoc](https://pkg.go.dev/badge/github.com/BrenekH/logange)](https://pkg.go.dev/github.com/BrenekH/logange)

Logging library which mimics Python's native `logging` library.

> :warning: Logange is an experimental and non-production ready library. Use at your own risk.

## Installation

Logange can be installed using the `go get -u github.com/BrenekH/logange` command.

## Concepts

Logange uses loggers, formatters, and handlers.

Loggers are part of a family tree with `logange.RootLogger` as the base node.
Logs will traverse the tree calling any attached handlers on the way to the root logger.

Handlers take log messages and send them to defined sources.
There are two handlers already available as part of Logange, `StdoutHandler` and `FileHandler` which are instantiated using `NewStdoutHandler` and `NewFileHandler` respectively.

Formatters are used by handlers to add additional information to the logs.

## Usage

### Basic Setup

The following code creates a logger called `main`, defines how the message should be formatted, and uses the `logange.StdoutHandler` to output messages to the console.

```go
package main

import "github.com/BrenekH/logange"

func main() {
    logger := logange.NewLogger("main")

    formatter := logange.StandardFormatter{FormatString: "${datetime}|${name}|${lineno}|${levelname}|${message}\n"}
    stdoutHandler := logange.NewStdoutHandler()

    stdoutHandler.SetFormatter(formatter)
    stdoutHandler.SetLevel(logange.LevelInfo)

    logger.AddHandler(&stdoutHandler)

    logger.Info("Hello, World!")
}
```
