# `plugin-kit-ai` npm package

Official `public-beta` npm wrapper for the `plugin-kit-ai` CLI itself.

Install paths:

```bash
npm i -g plugin-kit-ai
plugin-kit-ai version
```

```bash
npx plugin-kit-ai@latest version
```

This package downloads the matching published GitHub Release binary from `777genius/plugin-kit-ai`, verifies `checksums.txt`, and then runs the installed binary. It does not build Go from source and it does not widen `plugin-kit-ai install`.
