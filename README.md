# zap-plus

A lightweight and extensible logging library built on top of [Uber's Zap](https://github.com/uber-go/zap), providing a simplified interface for structured logging in Go applications.

## Features

- 🚀 **High Performance**: Built on Uber Zap's high-performance logging foundation
- 🎯 **Simple Interface**: Clean and intuitive API for common logging operations
- 🔧 **Configurable**: Support for different log levels and configurations
- 📦 **Lightweight**: Minimal overhead with focused functionality
- 🏗️ **Extensible**: Modular design for easy customization and extension

## Installation

```bash
go get github.com/hanzhang418/zap-plus
```

## Quick Start

```go
package main

import (
    "github.com/hanzhang418/zap-plus"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func main() {
    // Create a new logger with Info level
    logger, err := zap_plus.New(zapcore.InfoLevel)
    if err != nil {
        panic(err)
    }

    // Log some messages
    logger.Info("Application started", zap.String("version", "1.0.0"))
    logger.Warn("This is a warning", zap.Int("code", 404))
    logger.Error("An error occurred", zap.Error(err))
}
```

## API Reference

### Creating a Logger

```go
func New(level zapcore.Level) (logger.Logger, error)
```

Creates a new logger instance with the specified log level.

**Parameters:**
- `level`: The minimum log level (e.g., `zapcore.DebugLevel`, `zapcore.InfoLevel`, `zapcore.WarnLevel`, `zapcore.ErrorLevel`)

**Returns:**
- `logger.Logger`: Logger interface instance
- `error`: Error if logger creation fails

### Logger Interface

The `logger.Logger` interface provides the following methods:

```go
type Logger interface {
    Info(msg string, fields ...zap.Field)
    Warn(msg string, fields ...zap.Field)
    Error(msg string, fields ...zap.Field)
}
```

#### Methods

- **`Info(msg string, fields ...zap.Field)`**: Log an info-level message
- **`Warn(msg string, fields ...zap.Field)`**: Log a warning-level message
- **`Error(msg string, fields ...zap.Field)`**: Log an error-level message

All methods accept a message string and optional structured fields using Zap's field types.

## Examples

### Basic Logging

```go
logger, _ := zap_plus.New(zapcore.InfoLevel)

logger.Info("User logged in",
    zap.String("username", "john_doe"),
    zap.String("ip", "192.168.1.1"))

logger.Warn("High memory usage",
    zap.Float64("usage_percent", 85.5))

logger.Error("Database connection failed",
    zap.String("database", "users"),
    zap.Duration("timeout", time.Second*30))
```

### Different Log Levels

```go
// Debug level logger (logs everything)
debugLogger, _ := zap_plus.New(zapcore.DebugLevel)

// Error level logger (only logs errors)
errorLogger, _ := zap_plus.New(zapcore.ErrorLevel)

// Info level logger (logs info, warn, and error)
infoLogger, _ := zap_plus.New(zapcore.InfoLevel)
```

### Structured Logging with Fields

```go
logger, _ := zap_plus.New(zapcore.InfoLevel)

// Using various field types
logger.Info("Request processed",
    zap.String("method", "GET"),
    zap.String("path", "/api/users"),
    zap.Int("status_code", 200),
    zap.Duration("response_time", time.Millisecond*150),
    zap.Bool("cached", true))
```

## Project Structure

```
zap-plus/
├── zapplus.go              # Main package entry point
├── pkg/
│   └── logger/
│       └── logger.go       # Logger interface definition
├── internal/
│   └── builder/
│       └── builder.go      # Logger builder implementation
├── go.mod                  # Go module definition
├── LICENSE                 # MIT License
└── README.md              # This file
```

## Dependencies

- [go.uber.org/zap](https://github.com/uber-go/zap) - High-performance logging library
- [go.uber.org/multierr](https://github.com/uber-go/multierr) - Error handling utilities

## Requirements

- Go 1.23.4 or later

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built on top of [Uber's Zap](https://github.com/uber-go/zap) logging library
- Inspired by the need for a simpler logging interface while maintaining high performance