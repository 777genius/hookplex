# Publish JSON Contract

`plugin-kit-ai publish --format json` is the automation-facing report for the bounded top-level publish workflow.

## Stability

- Contract id: `plugin-kit-ai/publish-report`
- Current schema version: `1`
- Stability tier: public contract for CI and tooling

## Envelope

Every JSON report includes:

- `format`: always `plugin-kit-ai/publish-report`
- `schema_version`: currently `1`
- `channel`
- `target`
- `mode`
- `workflow_class`
- `detail_count`
- `details`
- `next_step_count`
- `next_steps`

Optional fields:

- `dest`: present for local marketplace-root workflows
- `package_root`: present for local marketplace-root workflows

## Workflow Classes

- `local_marketplace_root`: local Codex or Claude marketplace-root materialization flow
- `repository_release_plan`: Gemini repository or release publication planning flow

## Array and Map Guarantees

The following fields are always present in schema version `1`:

- `details`
- `next_steps`

`details` is always an object and `next_steps` is always an array, never `null`.

## Compatibility Rules

- Additive fields may appear in future schema versions
- Breaking changes require a new `schema_version`
- Consumers should branch first on `format`, then on `schema_version`
- Consumers should treat `workflow_class` as the primary dispatch field for channel-family automation
