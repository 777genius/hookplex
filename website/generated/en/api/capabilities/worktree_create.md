---
title: "worktree_create"
description: "Capability reference for worktree_create"
canonicalId: "capability:worktree_create"
surface: "capabilities"
section: "api"
locale: "en"
generated: true
editLink: false
stability: "public-beta"
maturity: "beta"
sourceRef: "docs/generated/support_matrix.md"
translationRequired: false
---
<DocMetaCard surface="capabilities" stability="public-beta" maturity="beta" source-ref="docs/generated/support_matrix.md" source-href="https://github.com/777genius/plugin-kit-ai/blob/main/docs/generated/support_matrix.md" />

# worktree_create

This capability shows one behavior across platforms. Open it when the action itself matters more than reading each platform tree separately.

## At A Glance

- Platforms with this capability: 1
- Related events: 1
- Current maturity: beta

## Read This With

- [Platform Behavior Model](/en/concepts/platform-behavior-model) when you need to decide whether capability-first is the right view.
- [Support Promise By Path](/en/reference/support-promise-by-path) when promise strength and operational cost matter more than the capability label alone.

## Platforms

- [`claude`](/en/api/platform-events/claude)

## Related Runtime Events

- `claude/WorktreeCreate`

## Coverage Table

| Platform | Event | Maturity | Contract |
| --- | --- | --- | --- |
| claude | WorktreeCreate | beta | runtime-supported but not stable |
