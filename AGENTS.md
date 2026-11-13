# tcpsystems/hetzner-mcp — Contributor & Agent Guide

`hetzner-mcp` (`cmd/mcp-hetzner`) is the Model Context Protocol (MCP) server for managing a Hetzner Cloud project on behalf of connected agents without exposing `HCLOUD_TOKEN` in agent contexts.

## Operating Rules

1. **Credential Isolation**:
   - `HCLOUD_TOKEN` (or `~/.config/hcloud/cli.toml` via `HCLOUD_CONTEXT`) is resolved once at process startup and never echoed to tool callers, logs, or error strings.
   - Streamable HTTP mode (`--transport http`) refuses non-loopback bind addresses unless `MCP_AUTH_TOKEN` is set (`cmd/mcp-hetzner/guard.go`).
2. **Quality Gate**:
   - Run `go vet ./...`, `go test -race ./...`, `staticcheck ./...`, and `govulncheck ./...` before every commit.
   - Keep `gofmt` clean across `cmd/` and `internal/`.
3. **Hetzner Safety Policy**:
   - Never run port scans, abuse-triggering probes, or unauthorized network traffic from Hetzner Cloud nodes.
4. **Branch & Identity Model**:
   - Default branch is `develop`; release branch is `production`.
   - Commit as `Petar Markov <petar@petarmarkov.co.uk>` using Conventional Commits.
