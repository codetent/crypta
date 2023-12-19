<div align="center">
  <br>
  <img src="docs/static/img/logo.svg" width="100" /><br>
  
  # Crypta
  
  Local developer credentials as simple as it should be.
  <br/><br/>
</div>

## What is crypta?

Crypta is a CLI tool that caches your local development credentials allowing you to enter them once and using it multiple times without beeing asked.

It can be also integrated in CI environments where a secret provider is available. In this case, the values can be provided using environment variables without adapting your scripts.

## Why do I need it?

If your build environment is pulling files from access protected services, you usually need credentials to authenticate. The best practice for this use case is to use your credentials instead of hardcoding them in your scripts.

Especially if you run the environment locally or even in the cloud without using any secret provider, you have to provide the credentials somehow to your build environment.

This is where crypta steps in: it is simply called in your build scripts when you need credentials to authenticate. If credentials are already cached, it will provide them. Otherwise, it will ask for them and cache them afterwards.

## Documentation

- CLI reference: [docs/pages](docs/pages/crypta.md)

## Installation

You will find the latest release of crypta in GitHub Releases section.

Otherwise you can install it from source using `go install`:
```sh
go install github.com/codetent/crypta@latest
```
