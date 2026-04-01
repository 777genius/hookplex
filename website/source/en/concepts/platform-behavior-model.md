---
title: "Platform Behavior Model"
description: "How to think about platform-first and capability-first views without confusing them with target choice or support strength."
canonicalId: "page:concepts:platform-behavior-model"
section: "concepts"
locale: "en"
generated: false
translationRequired: true
---

# Platform Behavior Model

Use this page when you already understand targets in general, but now need the clean mental model for reading `Platform Events` and `Capabilities`.

## Quick Rule

- use a **platform-first view** when one platform is already chosen and you need its exact events
- use a **capability-first view** when you want to compare similar behavior across platforms
- do not use either one to choose the target family itself; choose the target first

## Two Ways To Read The Same Layer

`plugin-kit-ai` exposes platform behavior through two complementary views:

- `Platform Events`: grouped by platform, then by event
- `Capabilities`: grouped by behavior, then by platform coverage

Neither view is “more correct.” They answer different questions.

## Platform-First View

Start platform-first when your main question sounds like:

- “What can Claude emit here?”
- “Which events exist for Codex?”
- “Which events on this platform are stable versus beta?”

This is the right view when one target platform is already known and you need the exact event surface.

## Capability-First View

Start capability-first when your main question sounds like:

- “Which platforms support notify-like behavior?”
- “Is this behavior cross-platform or platform-specific?”
- “How many platforms expose this capability at all?”

This is the right view when behavior comparison matters more than one platform tree.

## What This Layer Is Not

- It is not the same thing as choosing a target family.
- It does not replace the support promise by path.
- It does not mean Claude and Codex promise equal depth or equal stability for similar behavior.

The platform layer describes behavior surfaces, not equal support guarantees.

## The Safe Reading Order

1. Choose the target family.
2. Understand the support promise by path.
3. Then open the platform behavior layer through either:
   - `Platform Events` for platform-first reading
   - `Capabilities` for behavior-first reading

## Common Mistake

People often treat one shared capability name as proof of equal support across platforms. That is too strong.

A capability name tells you there is related behavior. It does **not** automatically mean:

- the same maturity
- the same event depth
- the same operational promise

## Read This With

- [Choose A Target](/en/guide/choose-a-target) when you still need the right target family
- [Support Promise By Path](/en/reference/support-promise-by-path) when you need the promise and cost framing
- [Platform Events](/en/api/platform-events/) when the platform is already known
- [Capabilities](/en/api/capabilities/) when behavior comparison is the main job

## Final Rule

Use `Platform Events` to answer “what exists on this platform?” and `Capabilities` to answer “where does this behavior exist across platforms?”.
