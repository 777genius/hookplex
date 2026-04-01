---
title: "API"
description: "Generated API reference for plugin-kit-ai."
canonicalId: "page:api:index"
section: "api"
locale: "en"
generated: false
translationRequired: true
aside: false
outline: false
---

<div class="docs-hero docs-hero--compact">
  <p class="docs-kicker">GENERATED REFERENCE</p>
  <h1>API Surfaces</h1>
  <p class="docs-lead">
    This reference is generated from the real CLI, packages, and structured metadata. It is split by
    public area so that API discovery stays predictable as the project grows.
  </p>
</div>

<div class="docs-grid">
  <a class="docs-card" href="./cli/">
    <h2>CLI</h2>
    <p>Commands exported from the live Cobra tree.</p>
  </a>
  <a class="docs-card" href="./go-sdk/">
    <h2>Go SDK</h2>
    <p>Public Go packages for stable integration paths.</p>
  </a>
  <a class="docs-card" href="./runtime-node/">
    <h2>Node Runtime</h2>
    <p>Typed runtime helpers for JS and TS consumers.</p>
  </a>
  <a class="docs-card" href="./runtime-python/">
    <h2>Python Runtime</h2>
    <p>Public Python runtime helpers only, not install wrappers.</p>
  </a>
  <a class="docs-card" href="./platform-events/">
    <h2>Platform Events</h2>
    <p>Event surfaces grouped by target platform.</p>
  </a>
  <a class="docs-card" href="./capabilities/">
    <h2>Capabilities</h2>
    <p>Capability-oriented view across platforms and events.</p>
  </a>
</div>

## Open The Right Surface

- Open `CLI` when you need commands, flags, or the authored workflow.
- Open `Go SDK` when you are building the strongest production-oriented runtime path.
- Open `Node Runtime` or `Python Runtime` when you need helper APIs for supported local runtime projects.
- Open `Platform Events` when you are choosing target-specific events.
- Open `Capabilities` when you want a cross-platform view of what a plugin can react to or enforce.

## What This API Section Covers

- the live Cobra command tree
- public Go packages
- shared runtime helper APIs for Node and Python
- platform-specific events
- capability-level cross-platform metadata
