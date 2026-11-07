#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

go mod download

printf '\nDependencies ready! Run `go run server.go`, then open port 8000 in the Ports tab.\n'
