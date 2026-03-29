#!/usr/bin/env python3
import json
import sys


def handle_notify():
    if len(sys.argv) < 3:
        sys.stderr.write("missing notify payload\n")
        return 1
    payload = json.loads(sys.argv[2])
    _ = payload
    return 0


def main():
    if len(sys.argv) < 2 or sys.argv[1] != "notify":
        sys.stderr.write("usage: main.py notify <json-payload>\n")
        return 1
    return handle_notify()


if __name__ == "__main__":
    raise SystemExit(main())
