<div align="center">
  <br>
  <img src="docs/logo.svg" width="100" /><br>

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

## Usage in CI environments

CI environments usually provide some sort of secret provider. It is common that secrets are made available to the
execution environment via environment variables. Crypta supports this usecase as well by pre-populating the secret store
with secret content from environment variables with the prefix `CRYPTA_SECRET_<KEY>`. The environment variables are
parsed when the crypta daemon is started.

In your build scripts, whenever you need to retrieve a secret, you would call `crypta get <KEY>`. If the build script is
executed in a CI environment, where a secret provider is available, the secret is pre-populated via the
`CRYPTA_SECRET_<KEY>` environment variable. If the build script is executed in a local development environment on the
other hand, a secret provider is usually not available. Therefore, no pre-population of the secret store takes place and
Crypta would now request the user to input the value for the secret (e.g., their user credentials). This
allows you to use the same build script in your CI and local development environment, without any adaptations or special 
case handling required to handle the secrets.
