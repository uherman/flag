# flag

A CLI tool that prints country flags as colored ASCII art in your terminal.

Uses 24-bit true color ANSI escape codes and Unicode half-block characters for high-quality rendering. Supports 193 countries.

## Installation

```sh
go install github.com/uherman/flag@latest
```

## Usage

```sh
# Print a specific flag
flag -sweden
flag -usa
flag -jp        # aliases work too

# Print all flags
flag -all

# List available flags
flag -list
```

## Requirements

- A terminal with true color (24-bit) support
- Go 1.22+ (for building from source)

## Examples

```
$ flag -norway
```
```
$ flag -brazil
```
```
$ flag -japan
```
