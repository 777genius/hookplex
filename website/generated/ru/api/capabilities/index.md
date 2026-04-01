---
title: "Capabilities"
description: "Generated capability reference"
canonicalId: "page:api:capabilities:index"
surface: "capabilities"
section: "api"
locale: "ru"
generated: true
editLink: false
stability: "public-beta"
maturity: "beta"
sourceRef: "docs/generated/support_matrix.md"
translationRequired: false
---
# Capabilities

Capabilities дают cross-platform view на runtime behavior, а не только platform tree.

- Открывайте эту зону, когда хотите понять не platform name, а само действие или реакцию.
- Это лучший вход, если вы сравниваете похожее поведение между Claude и Codex.

## С чего лучше начать

- Нужна карта поведения между платформами: начинайте отсюда.
- Нужна точная event-level детализация по одной платформе: переходите в `Platform Events`.

## Когда не нужно начинать отсюда

- Если вы ещё не понимаете сами target families, сначала прочитайте `/guide/what-you-can-build` и `/guide/choose-a-target`.

## С чем читать вместе

- [Модель поведения платформ](/ru/concepts/platform-behavior-model), если нужно понять, когда capability-first view сильнее platform-first.
- [Обещания поддержки по путям](/ru/reference/support-promise-by-path), если важнее support promise, чем сама capability.

## Карта capabilities

| Capability | Платформы | Events |
| --- | --- | --- |
| config_change | 1 | 1 |
| notify | 2 | 2 |
| permission_request | 1 | 1 |
| post_tool | 1 | 1 |
| post_tool_failure | 1 | 1 |
| pre_compact | 1 | 1 |
| prompt_submit_gate | 1 | 1 |
| session_end | 1 | 1 |
| session_start | 1 | 1 |
| setup | 1 | 1 |
| stop_gate | 1 | 1 |
| subagent_start | 1 | 1 |
| subagent_stop | 1 | 1 |
| task_completed | 1 | 1 |
| teammate_idle | 1 | 1 |
| tool_gate | 1 | 1 |
| worktree_create | 1 | 1 |
| worktree_remove | 1 | 1 |

## Переходите по capabilities

- [`config_change`](/ru/api/capabilities/config_change)
- [`notify`](/ru/api/capabilities/notify)
- [`permission_request`](/ru/api/capabilities/permission_request)
- [`post_tool`](/ru/api/capabilities/post_tool)
- [`post_tool_failure`](/ru/api/capabilities/post_tool_failure)
- [`pre_compact`](/ru/api/capabilities/pre_compact)
- [`prompt_submit_gate`](/ru/api/capabilities/prompt_submit_gate)
- [`session_end`](/ru/api/capabilities/session_end)
- [`session_start`](/ru/api/capabilities/session_start)
- [`setup`](/ru/api/capabilities/setup)
- [`stop_gate`](/ru/api/capabilities/stop_gate)
- [`subagent_start`](/ru/api/capabilities/subagent_start)
- [`subagent_stop`](/ru/api/capabilities/subagent_stop)
- [`task_completed`](/ru/api/capabilities/task_completed)
- [`teammate_idle`](/ru/api/capabilities/teammate_idle)
- [`tool_gate`](/ru/api/capabilities/tool_gate)
- [`worktree_create`](/ru/api/capabilities/worktree_create)
- [`worktree_remove`](/ru/api/capabilities/worktree_remove)
