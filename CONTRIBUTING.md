# Contributing

This guide describes all necessary steps for contributing to crypta.

## Setup

The following prerequisites are required for working with the project:

**Required:**
- Docker
- Visual Studio Code + Dev Containers extension

**Recommended:**
- mutagen-compose (mutagen is used for file sharing replacement which is more performant than the native docker share)
  - Visual Studio Code settings: set the path of the mutagen-compose binary at
    `Dev > Containers: Docker Compose Path`

Devcontainers are used for the project. The provided container includes the necessary environment for building, testing
and linting. Additionally, it is used for the GitHub action workflows.

To start development, the repository has to be cloned first:

```
git clone https://github.com/codetent/crypta.git
```

Afterwards, it can be opened up using Visual Studio Code.
The container can then be started by pressing `F1` and calling `Dev Containers: Rebuild and Reopen in Container`.

> [!NOTE]  
> If mutagen-compose is not used, the repository has to be cloned again inside the dev container at /workspace.

## Generate

The connection between client and daemon is realized using [connect-rpc](https://connectrpc.com/). It uses protobuf
for the communication where the protocol is defined using `.proto` files.

Using the `buf` CLI a corresponding server and client implementation is generated based on the given `.proto` files.
Each time these files are changed, the go files in the `gen` folder have to be recreated by calling:

```
buf generate
```

Additionally, these files can be statically analyzed (which is done in the workflows automatically) using:

```
buf lint
```

## Lint

The project is linted using `golangci-lint`.
It can be run by calling:

```
golangci-lint run
```

## Build

During local development, the CLI can be executed using go directly:

```
go run main.go
```

For getting the compiled binary, `go build` can be utilized.

In the workflows, several binaries for different operating systems are built using `goreleaser`.
The workflow can be tested, if needed:

```
goreleaser --clean --snapshot
```

## Test

The project uses two types of tests for its verification:

- Unit tests are used to verify single isolated units while mocking other units.
The tests are defined at module level, by creating a file with the module name and a `_test` suffix.
- End-to-end tests are used for verifying the functionality and performance of the entire application from start to
finish from the user's perspective. They are located in the the `test/e2e` folder.

To run the tests (unit & e2e), simply call:

```
ginkgo ./...
```

For updating mocks, execute:

```
mockery
```

## Documentation

The documentation of the project is hosted using GitHub pages.
For that, a static website is built based on Markdown files located in `docs/pages`.

The CLI commands are documented directly in the code. Afterwards, a reference is generated by extracting the
necessary information into Markdown files (`docs/pages/cli`). Each time the CLI commands are updated (e.g., help text,
flags), the following command has to be called to regenerate the reference:

```
./hack/gen_docs.sh
```

## Workflows

Following workflows are available via GitHub actions:

### Merge Verification

For each pull request that is going to be merged to the `main` branch, a merge verification is automatically executed.
The following steps are executed:

1. Run statical analysis
2. Execute unit tests
3. Build binaries as snapshots
4. Build documentation

#### Labels

The following pull request labels shall be used for determining the change type and consequently the next version:
- `breaking`: Changes breaking the user interface or existing functionality
- `feature`: Changes adding new features without breaking existing functionality
- `bug`: Patches for bugs for existing functionality
- `chore`: No production code change

### Branch Verification

After changes are merged to the `main` branch, a branch build is triggered, which is based on the merge verification,
except it:

- Builds binaries with next version
- Creates draft release with changelog and binaries
- Deploys documentation to GitHub pages

### Docker Image Build

To build a new version of the docker image used locally and for the workflows, the manual workflow
`Build and push dev image` has to be executed with the commit which should be built as input.

It builds the docker image from the provided commit and pushes it to the registry.
Afterwards, the used tag has to be updated in all files referencing `codetent/crypta-dev` (use search).

### Release

Since crypta follows continuous delivery methodologies for its builds, releasing a new version is a non-event.
Instead just an existing draft release must be published.

This is done by going to the `Releases` section, selecting the draft, editing it, checking if it is complete and
publishing it by clicking on `Publish release`.
