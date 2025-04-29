# ulifix

A simple CLI tool to generate ULIDs with an optional prefix or suffix.

## Usage

```bash
# Generate a ULID with a prefix
ulifix myprefix_

# Generate a ULID with a suffix
ulifix -s _mysuffix

# Generate only a ULID
ulifix
```

## Installation

### Install via go install (recommended)

```bash
go install github.com/mkusaka/ulifix@latest
```

> Make sure your $GOPATH/bin (or $GOBIN) is in your $PATH.

### Manual build

```bash
# Clone the repository
git clone https://github.com/mkusaka/ulifix.git
cd ulifix

# Build the binary
go build

# Install the binary (optional, requires setting up GOPATH/bin in your PATH)
go install
```

## Development

Requires Go 1.18 or later.

```bash
# Get dependencies
go get

# Run tests
go test
```

