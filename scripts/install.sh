#!/usr/bin/env sh
# Bootstrap: fetch a released hookplex CLI tarball (when you publish releases).
# After `hookplex` is on PATH, install third-party plugin binaries with:
#   hookplex install owner/repo --tag vX.Y.Z   OR   --latest
#   [--dir bin] [--force] [--pre] [--output-name NAME]
# (see root README — checksums.txt + .tar.gz or raw *-GOOS-GOARCH[.exe].)
#
# Usage: VERSION=v0.1.0 ./scripts/install.sh
# Override: HOOKPLEX_DOWNLOAD_URL to fetch a custom tarball.

set -e

VERSION="${VERSION:-}"
if [ -z "$VERSION" ]; then
  echo "Set VERSION (e.g. VERSION=v0.1.0) to download a release build." >&2
  echo "Until releases are published, build from source: make build-hookplex" >&2
  exit 1
fi

if [ -n "$HOOKPLEX_DOWNLOAD_URL" ]; then
  url="$HOOKPLEX_DOWNLOAD_URL"
else
  echo "HOOKPLEX_DOWNLOAD_URL is not set; set it to the release archive URL for your OS/arch." >&2
  exit 1
fi

tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT
curl -fsSL "$url" -o "$tmp/archive.tar.gz"
tar -xzf "$tmp/archive.tar.gz" -C "$tmp"
echo "Extracted to $tmp — copy the hookplex binary to your PATH." >&2
