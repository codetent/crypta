#!/bin/sh

go run docs/hack/cli/main.go

cd docs
npm ci --no-audit --no-fund
npm run build
