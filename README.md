keyring-defaults
================
[![CI](https://github.com/lox/keyring-defaults/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/lox/keyring-defaults/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/lox/keyring-defaults.svg)](https://pkg.go.dev/github.com/lox/keyring-defaults)

Default first-party provider list for [`github.com/lox/keyring/v2`](https://github.com/lox/keyring).

This module exists so applications can opt in to the common desktop and
command-backed providers without adding those imports to the lightweight root
keyring module.

## Usage

```bash
go get github.com/lox/keyring-defaults
```

```go
import (
	"context"

	defaults "github.com/lox/keyring-defaults"
	"github.com/lox/keyring/v2"
)

ctx := context.Background()

ring, err := keyring.Open(ctx,
	keyring.WithServiceName("example"),
	keyring.WithProviders(defaults.Providers()...),
)
```

Pass application-owned fallbacks as extra providers:

```go
providers := defaults.Providers(
	keyring.FileProvider(
		keyring.FileDir("/path/to/keyring"),
		keyring.FilePrompt(keyring.FixedStringPrompt("passphrase")),
	),
)
```

Applications that need explicit macOS Keychain trust policy can override the
default keychain provider without changing its position:

```go
providers := defaults.Providers(
	defaults.KeychainTrustApplication(false),
	keyring.FileProvider(
		keyring.FileDir("/path/to/keyring"),
		keyring.FilePrompt(keyring.FixedStringPrompt("passphrase")),
	),
)
```

`Providers` returns providers in this order:

1. Windows Credential Manager
2. macOS Keychain
3. Secret Service
4. KWallet
5. Linux keyctl
6. pass
7. any extra providers passed by the application

Unsupported platform providers return `keyring.ErrUnavailable` when opened, so
`keyring.Open` can fall through to the next provider.

The encrypted file provider is intentionally not included by default because its
directory and prompt are application policy. Pass `keyring.FileProvider(...)` to
`Providers` when you want it as a fallback.

## Development

```bash
go test ./...
go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.12.2 run
```
