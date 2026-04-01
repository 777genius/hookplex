---
title: "Reference Repo Strategy"
description: "How to choose, protect, and update one reference repo before a team turns a path into templates, defaults, or rollout guidance."
canonicalId: "page:guide:reference-repo-strategy"
section: "guide"
locale: "en"
generated: false
translationRequired: true
---

# Reference Repo Strategy

Use this page when the team already agrees that one repo should become the public or internal reference point for a path, but you want that repo to teach the right standard instead of accidental folklore.

## Choose In 60 Seconds

- use this page when one repo is about to become the example other repos copy
- use this page before updating templates, starter guidance, or team rollout docs
- use this page when several repos exist, but nobody agrees which one should define the baseline
- do not use this page when you still have not chosen the path itself; fix that first

## What This Page Helps You Decide

- which repo should become the reference repo
- what that repo must prove before the team copies it
- how to update a reference repo without turning it into a snowflake

## What A Reference Repo Is

A reference repo is not just the first repo that happened to work.

It is the repo that:

- best represents the intended product shape
- passes the full public contract cleanly
- is visible enough that the team can actually learn from it
- stays close enough to the standard path that its quirks do not become fake rules

## What A Reference Repo Is Not

- not the most complex repo in the estate
- not the repo with the most custom exceptions
- not the repo one maintainer understands privately
- not a place to prove every possible target, runtime, or delivery mode at once

## Choose A Good Reference Repo

Pick the repo that is:

- active enough to matter
- simple enough to explain
- representative of the path you want others to copy
- already close to passing `doctor -> render -> validate --strict`

If no repo fits those conditions, do not force one. Create or clean one reference repo first.

## What The Reference Repo Must Prove

Before the team treats a repo as the standard, it should prove:

- the chosen target and runtime are explicit
- rendered outputs are reproducible
- CI runs the same readiness contract as local development
- README or repo docs explain the path without private tribal knowledge
- release notes and support promises linked by the repo still match the chosen path

## The Safe Lifecycle

1. Pick one path first.
   A reference repo cannot compensate for an unclear target, runtime, or delivery model.
2. Clean one repo.
   Make it pass the full contract without hand-patched generated files.
3. Link the public baseline.
   Point to the current release and support rule behind that path.
4. Let other repos copy it deliberately.
   Only after the repo is clean should starter guidance, templates, or rollout tracking point to it.
5. Keep it boring.
   A reference repo should evolve carefully, not become the most experimental repo in the estate.

## When To Replace The Reference Repo

Replace it when:

- the repo has accumulated path-specific exceptions
- the team now uses a different mainstream path
- the repo no longer represents the cleanest public baseline

Do not replace it just because a newer repo exists. Replace it when the current reference no longer teaches the right standard.

## What Not To Do

- do not let the most advanced repo become the standard by accident
- do not upgrade the reference repo in ways you cannot explain publicly
- do not point several conflicting repos to the team as “examples”
- do not let release notes, CI, and repo docs disagree about the chosen path

## Best First Stops

- Need the team-level adoption path: [Team Adoption](/en/guide/team-adoption)
- Need the multi-repo rollout path: [Team-Scale Rollout](/en/guide/team-scale-rollout)
- Need safe correction before standardization: [Path Recovery](/en/guide/path-recovery)
- Need the repo contract itself: [Repository Standard](/en/reference/repository-standard)
- Need the version and support baseline: [Version And Compatibility Policy](/en/reference/version-and-compatibility) and [Support Promise By Path](/en/reference/support-promise-by-path)

## Final Rule

If another maintainer cannot look at the reference repo and learn the intended standard without private explanation, it is not a real reference repo yet.
