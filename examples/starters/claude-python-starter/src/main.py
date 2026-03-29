#!/usr/bin/env python3
import json
import sys


def read_stdin_json():
    return json.load(sys.stdin)


def handle_claude(hook_name):
    if hook_name not in {"Stop", "PreToolUse", "UserPromptSubmit"}:
        sys.stderr.write("usage: main.py <Stop|PreToolUse|UserPromptSubmit>\n")
        return 1
    event = read_stdin_json()
    _ = event
    sys.stdout.write("{}")
    return 0

def main():
    if len(sys.argv) < 2:
        sys.stderr.write("usage: main.py <Stop|PreToolUse|UserPromptSubmit>\n")
        return 1

    hook_name = sys.argv[1]
    return handle_claude(hook_name)


if __name__ == "__main__":
    raise SystemExit(main())
