# Tagger

## Introduction

Tag Git/Github repos and Docker images all with the same value

## Contents

- [Install](#install)
- [Usage](#usage)
- [Defaults](#defaults)
- [TODO](#todo)

## Install

Install via `brew`:

```bash
brew tap veritone/veritone
brew install tagger
```

Or `curl` from releases:

```bash
curl -o tagger https://github.com/veritone/tagger/releases/download/1.0.0/tagger-darwin-amd64
sudo mv tagger /usr/bin
```

## Usage

Create a `tagger.yml` file:

```yml
git:
  - dir: $GOPATH/src/github.com/veritone.com/veritone-sdk
  - dir: $GOPATH/src/github.com/veritone.com/veritone-sample-app-react

docker:
  - from_image: veritone/aiware
  - from_image: veritone/logger
```

Run `tagger`:

```bash
tagger 0.1.0
```

## Defaults

`tagger.yml` files have option fields for each group.

### Git

| Key | Value |
| --- | --- | 
| `ref` | `master` |
| `remote` | `origin` |
| `tag` | CLI Tag Value |

A `tagger.yml` configuration like:

```yml
git:
  - dir: $GOPATH/src/github.com/veritone/veritone-sdk
```

Will be filled in to look like:

```yml
git:
  - dir: $GOPATH/src/github.com/veritone/veritone-sdk
    ref: master
    remote: origin
    tag: <CLI_TAG_VALUE>
```

### Docker

| Key | Value |
| --- | --- | 
| `from_tag` | `latest` |
| `to_image` | Value from `from_image` |
| `to_tag` | CLI Tag Value |

A `tagger.yml` configuration like:

```yml
docker:
  - from_image: bash
```

Will be filled in to look like:

```yml
git:
  - from_image: bash
    from_tag: latest
    to_image: bash
    tag: <CLI_TAG_VALUE>
```

## TODO

- [ ] Batch logging for better stdout logs
