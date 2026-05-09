# Contributing

This document covers local development for plugin authors. For user-facing
plugin docs (configuration, supported resources, examples), see
[README.md](README.md).

## Prerequisites

- Go 1.25+
- [Pkl CLI](https://pkl-lang.org/main/current/pkl-cli/index.html)
- SFTP server for testing

## Local Installation

```bash
make install
```

## Building

```bash
make build      # Build plugin binary
make test       # Run unit tests
make lint       # Run linter
make install    # Build + install locally
```

## Local Testing

```bash
# Start test SFTP server
docker run -d --name sftp-test -p 2222:22 atmoz/sftp testuser:testpass:::upload

# Install plugin locally
make install

# Start formae agent
SFTP_USERNAME=testuser SFTP_PASSWORD=testpass formae agent start

# Apply example resources
formae apply --mode reconcile examples/basic/main.pkl

# Clean up
docker rm -f sftp-test
```

## Conformance Testing

Run the full CRUD lifecycle + discovery tests:

```bash
make conformance-test                  # Latest formae version
make conformance-test VERSION=0.80.0   # Specific version
```
