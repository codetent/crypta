#!/bin/sh

go run docs/hack/cli/main.go

cd docs
npm ci
npm run build
