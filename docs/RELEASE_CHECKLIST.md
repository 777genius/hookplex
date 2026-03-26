# Release Checklist

Use this checklist for any pre-`v1` beta release and for the future `v0.9` freeze milestone.

## Required Before Tagging

- `make test-required` green
- `make vet` green
- `make release-gate` may be used as the canonical local gate shortcut
- generated artifacts in sync
- support matrix matches shipped claims
- changelog updated
- support/status/release docs updated if contract changed
- candidate commit SHA recorded

## Extended / Live Recording

- `extended` workflow result recorded
- `live` workflow result recorded, or an explicit waiver is noted in release notes
- any skipped real-CLI smoke reason is written down
- waiver justification explicitly states why the failure is outside hookplex contract scope

## Beta-Breaking Changes

- migration note written when beta user code or scaffold output changes
- deprecation or removal called out in docs/changelog
- stable-candidate set impact reviewed
- [V0_9_AUDIT.md](./V0_9_AUDIT.md) updated when the declared `v1` candidate set changes
- [MIGRATIONS.md](./MIGRATIONS.md) updated for every beta-breaking user-visible change

## Rehearsal Completion

- each candidate surface is marked `stable-approved`, `stays-beta`, or `blocked`
- no core stable-set surface remains `blocked`
- release notes draft exists
- rehearsal worksheet exists
- known limitations are written down

## `v0.9` Freeze Check

- no new public-beta surfaces added unless required to finish the declared `v1` set
- remaining work limited to bug fixes, docs, migration, e2e hardening, and release tightening
