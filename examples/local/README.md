# Repo-Local Plugin Examples

These examples are reference implementations for the fast local plugin entrance layer.

- [codex-python-local](./codex-python-local): repo-local `codex-runtime` example for Python teams using `plugin-kit-ai bootstrap .`, `.venv`, `validate --strict`, and launcher-based `notify`
- [codex-node-local](./codex-node-local): repo-local `codex-runtime` example for Node teams using `plugin-kit-ai bootstrap .`, `validate --strict`, and launcher-based `notify`
- [codex-node-typescript-local](./codex-node-typescript-local): repo-local `codex-runtime` example for TypeScript teams using `plugin-kit-ai doctor .`, `plugin-kit-ai bootstrap .`, and built output under `dist/`

These Node/TypeScript and Python examples are the `public-stable` repo-local local-runtime subset.
Launcher-based `shell` authoring remains `public-beta` and is covered through runtime docs plus `polyglot-smoke`, not through a checked-in local example repo.
They complement, not replace, the production reference repos in [../plugins/README.md](../plugins/README.md).
Go remains the strongest production path; use the production examples when you want the clearest long-term support and release story.
