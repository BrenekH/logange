# Logange

Logging library which mimics Python's native `logging` library.

> :warning: Logange is an experimental and non-production ready library. Use at your own risk.

## Installation

Logange can be installed using the `go get -u github.com/BrenekH/logange` command.

## Usage

### Basic Setup

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
