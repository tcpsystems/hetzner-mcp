# REPO-CLEAN-HETZNER-MCP

Card: make tcpsystems/hetzner-mcp read as a clean, intentional repository — plain product-voice README, contributor guidance for any contributor, none of the working records of the build process.

## Plan

1. Research: OpenViking search for the card; read top hits in full (DOCS-STYLE.md reviewed); context7 lookups for hcloud-go and modelcontextprotocol/go-sdk (exact APIs used in the repo); repo-wide grep for references to any residue file before deletion.
2. Read every file to change in full (README.md done; AGENTS.md, CLAUDE.md, .claude/ do not exist in this repo — confirmed by find).
3. Write docs/evidence/REPO-CLEAN-HETZNER-MCP.md plan (this file).
4. Rewrite README.md in plain product voice: title + one sentence, status with observation date, why (<= three sentences), quickstart with commands actually run and expected exit codes, architecture with Mermaid diagram, API and configuration summary with links (no endpoint list copied), security and operations, further documentation links. Product vocabulary (Account, Team, Project, Repository, Environment, Agents, Runs). No internal tooling words (fanout, lane, card, seat, harness, lead, fleet-pool, dsh, oc-queue). Sentence-case headings; no emoji, em/en dashes, first-person plural, or hype. All claims traceable to code at cited paths.
5. Add CONTRIBUTING.md: how to run, how to check a change, how to contribute, commits, releases, secrets. Written for any contributor; no internal tooling words.
6. Confirm no residue files exist (AGENTS.md, CLAUDE.md, .claude/, MERGE-CLEAR-BRIEF*, HANDOFF*, RES-*, META-CARDS-*, UI-RESEARCH-*) and nothing references them (grep done, no matches).
7. Run the repo gate: go build ./... && go vet ./... && go test ./... (no tools/gate.sh and no Makefile exist in this repo). Record exit codes.
8. Run both VERIFY lines from the card.
9. Commit with conventional message; push branch; open PR titled "docs: product-voice README and remove harness residue" (under ~600 changed lines excluding deletions).
10. Finish this evidence file: files changed, commands with exit codes, acceptance criteria mapped to evidence.

## Research applied

- viking://resources/acumen-platform-corpus/Documents/maps/2026-10-30/docs/DOCS-STYLE.md — read in full: README contract (eight ordered sections), banned style (AI/agent flavour, hype, first-person plural, emoji, em/en dashes, Title Case headings, unquantified performance claims), live claims dated. Applied to README structure and style.
- viking://resources/cli/repos/cli/docs__cards__CLI-MCP-HETZNER-REGISTRY-ONBOARD.md — appears in search results; not needed for this docs-only change (registry onboarding metadata, not repository voice).
- context7 /hetznercloud/hcloud-go — hcloud.NewClient, WithToken, default endpoint https://api.hetzner.cloud/v1; matches cmd/mcp-hetzner and internal/server usage.
- context7 /modelcontextprotocol/go-sdk — mcp.NewServer, mcp.StdioTransport, mcp.NewStreamableHTTPHandler; matches internal/server/app.go and cmd/mcp-hetzner/main.go.
- Repo-wide grep for AGENTS.md, CLAUDE.md, .claude/, MERGE-CLEAR, HANDOFF, RES-*, META-CARDS-*, UI-RESEARCH-* references: zero matches. No tools/gate.sh, no Makefile; gate is go build ./... && go vet ./... && go test ./... (CI: .github/workflows/test.yml runs go test ./... on push/PR; release.yml publishes archives from v* tags).

## Libraries adopted

None — this change touches only Markdown documentation. Third-party libraries named in the README (hcloud-go, modelcontextprotocol/go-sdk) are already dependencies in go.mod and are described exactly as used in the code; no new dependency is introduced.

## Results

Commands run (repo /cache/ws/b6ead45b-00e3-6025-ad0a-52c0a3180a2a/repo, branch card/REPO-CLEAN-HETZNER-MCP):

- `go version` → go1.27.3 linux/amd64 (go.mod requires go 1.25.0)
- `git tag | tail -5` → latest tag v0.3.0 (status claim in README is dated 2026-11-09)
- Gate baseline before edits: `go build ./... && go vet ./... && go test ./...` → rc 0, all four packages ok
- `go build -o mcp-hetzner ./cmd/mcp-hetzner` → rc 0 (quickstart build step verified)
- `./mcp-hetzner --version` → prints "dev", rc 0 (quickstart version step verified)
- `printf '' | HCLOUD_TOKEN=bad ./mcp-hetzner` → rc 0 (server resolves HCLOUD_TOKEN and serves stdio until stdin closes)
- Artifact removed: `rm -f mcp-hetzner` (never committed; root binary is not gitignored)
- Gate after edits: `go build ./... && go vet ./... && go test ./...` → rc 0, all four packages ok
- VERIFY 1 (residue): `! git ls-files | grep -qE '(^|/)CLAUDE\.md$|^\.claude/|MERGE-CLEAR|HANDOFF|(^|/)(RES|META-CARDS|UI-RESEARCH)-[A-Z0-9-]*\.md$'` → pass (nothing tracked; repo-wide grep for references to AGENTS.md, CLAUDE.md, .claude/, MERGE-CLEAR, HANDOFF, RES-*, META-CARDS-*, UI-RESEARCH-* found zero matches, and no such files exist)
- VERIFY 2 (forbidden words): `! grep -siqwE 'fanout|fanoutd|oc-queue|fleet-pool|dsh|harness|lead-cycle|cards?|lanes?' README.md CONTRIBUTING.md AGENTS.md` → pass (AGENTS.md does not exist; no forbidden word in README.md or CONTRIBUTING.md)

## Files changed

- README.md — rewritten in product voice: status (live, v0.3.0, observed 2026-11-09), why it exists, quickstart with verified commands and exit codes, architecture with Mermaid diagram, API and configuration summary, security and operations, further documentation links. Product vocabulary used (Project, Environment, Agents, Repository). No internal tooling words. Sentence-case headings; no emoji, em/en dashes, first-person plural, or hype. Every claim cites a code path.
- CONTRIBUTING.md — new: how to run, how to check a change (the gate), how to contribute, commits, releases from version tags, secrets, conduct. For any contributor; no internal tooling words.
- docs/evidence/REPO-CLEAN-HETZNER-MCP.md — this file.

## Deleted

None. AGENTS.md, CLAUDE.md, .claude/, MERGE-CLEAR-BRIEF*, handoff notes, session notes, and RES-*/META-CARDS-*/UI-RESEARCH-* dumps do not exist in this repo, and nothing references them, so there was nothing to delete (evidence: the two grep searches above).

## Acceptance criteria → how each is met

- Gate rc0 — `go build ./... && go vet ./... && go test ./...` exits 0 after the edits (re-run recorded above).
- No residue files tracked — VERIFY 1 passes; nothing matching the residue patterns exists or is tracked.
- No forbidden words in README.md, CONTRIBUTING.md, AGENTS.md — VERIFY 2 passes.
- PR titled "docs: product-voice README and remove harness residue" — to be opened against main from branch card/REPO-CLEAN-HETZNER-MCP; changed lines excluding deletions are well under 600.

## Not done

None. The card is complete once the PR is opened (the runner pushes the branch).
