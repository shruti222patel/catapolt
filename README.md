# Catapault
Automates setting up different workspaces for different projects.

## Setup
1. Install [go](https://go.dev/dl/) -- v1.13.
2. `git clone` repo.
3. `cd` into project root.
4. `go mod download`

## Run
`go run main.go catapault` -- sets up workspace based on the `config.yaml` for this project.

## Release
1. Install `brew install goreleaser`
2. 