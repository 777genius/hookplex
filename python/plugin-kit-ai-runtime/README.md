# plugin-kit-ai-runtime (PyPI authoring helper)

Official Python helper package for launcher-based `plugin-kit-ai` plugins.

Use it when you want the supported handler-oriented API in a shared dependency instead of copying a local helper file into every repo.

Install:

```bash
pip install plugin-kit-ai-runtime
```

Example:

```python
from plugin_kit_ai_runtime import CodexApp, continue_

app = CodexApp()


@app.on_notify
def on_notify(event):
    _ = event
    return continue_()


raise SystemExit(app.run())
```

Notes:

- Go is still the recommended path when you want the most self-contained delivery model.
- Python authoring remains a stable supported lane, but the machine running the plugin still needs Python `3.10+`.
- The helper API mirrors the generated `src/plugin_runtime.py` scaffold surface.
