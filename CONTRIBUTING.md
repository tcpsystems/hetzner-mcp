# Contributing to hetzner-mcp

This repository is open to any contributor. Work happens on branches against main, and changes land through pull requests.

## How to run

```bash
go build -o mcp-hetzner ./cmd/mcp-hetzner
HCLOUD_TOKEN=... ./mcp-hetzner
```

Set HCLOUD_TOKEN in the environment, or keep the token in ~/.config/hcloud/cli.toml. The server binds to loopback for streamable HTTP unless MCP_AUTH_TOKEN is set.

## How to check a change

```bash
go build ./...
go vet ./...
go test ./...
```

All three must exit 0. This is the repository gate; there is no separate gate script. Run it before every push and fix failures until green.

## How to contribute

1. Create a branch from main.
2. Make one change at a time. Keep runbooks, ADRs, contracts, and anything the build or gate reads where they are.
3. Add or update tests for changed behaviour. Never skip or weaken a test to make it pass.
4. Run the check commands above until they are green.
5. Commit with a conventional commit message: `<type>(<scope>): <what>`. Types: build, ci, docs, feat, fix, refactor, style, test. Use the imperative mood and sentence case.
6. Open a pull request against main. Reference the issue that motivated the change.

## Releases

Releases are cut from version tags (v1.2.3). Pushing a tag runs tests, builds archives for linux, darwin, and windows on amd64 and arm64, and publishes them as a GitHub release. The linker stamps the version into internal/server/version.go.

## Secrets

Never commit tokens, keys, or connection strings. Configure the server through environment variables (HCLOUD_TOKEN, MCP_AUTH_TOKEN) or the hcloud CLI config. Secrets enter through the environment at runtime; they never belong in the repository.

## Conduct

Treat other contributors with respect. Ask before you assume.
