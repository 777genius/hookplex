---
title: "Target Support"
description: "Generated target support summary"
canonicalId: "page:reference:target-support"
surface: "reference"
section: "reference"
locale: "ru"
generated: true
editLink: false
stability: "public-stable"
maturity: "stable"
sourceRef: "docs/generated/target_support_matrix.md"
translationRequired: false
---
# Поддержка target’ов

Используйте эту страницу, когда нужно быстро понять, какой target production-ready, а какой остаётся packaging-only или workspace-config lane.

| Target | Production Class | Runtime Contract | Install Model |
| --- | --- | --- | --- |
| claude | production-ready | stable runtime subset | marketplace or local |
| codex-package | package lane | official package only | marketplace or local |
| codex-runtime | runtime lane | stable notify runtime | repo-local |
| gemini | packaging-only | packaging, not runtime | copy install |
| cursor | packaging-only | workspace-config lane | workspace config |
| opencode | packaging-only | workspace-config lane | workspace config |

Для полной framing-картины свяжите эту матрицу с [Границей поддержки](/ru/reference/support-boundary) и [Моделью target’ов](/ru/concepts/target-model).
