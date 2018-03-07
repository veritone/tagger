# Tagger

## Introduction

In the software industry, [semantic versioning](https://semver.org/) is a way of life to manage dependency hell with your applications and systems. In ideal worlds, your team will have mature git-flow and SDLC pipelines managed by CI/CD systems like Jenkins, Travis, Wercker, etc. However, with small teams or new projects, you may not have these mature pipelines but still need semantic versioning. Veritone's Tagger can satisfy this usecase. 

Tagger is distributed as a single binary and is managed via a single configuration file. Tagger handles tagging Git repos and updating Github & tagging Docker images and pushing to the appropriate registries.

## Installation

On OS X systems, you can use [Homebrew](https://brew.sh/) to install Tagger:

```bash
brew tap veritone/veritone
brew install tagger
```

## Usage 

Generate a `tagger.yml` file that will describe which Git repos & Docker images to handle:

```yml
git:
  - dir: ~/github.com/veritone.com/veritone-sample-app-react

docker:
  - from_image: veritone/logger
```

Then run `tagger` from the same directory and define which version you want everything to have:

```bash
tagger 0.1.0
```

Here's an explanation of the operations completed by the command above:

### Git

- Change directory to `~/github.com/veritone.com/veritone-sample-app-react`
- `git checkout master`
- `git tag 0.1.0`
- `git push origin 0.1.0`

### Docker

- `docker pull veritone/logger:latest`
- `docker tag veritone/logger:latest veritone/logger:0.1.0`
- `docker push veritone/logger:0.1.0`

## Overrides

It's possible to override the default values by adding the keys in the `tagger.yml` file. Here is a `tagger.yml` can look like:

```yml
git:
  - dir: $GOPATH/src/github.com/veritone.com/veritone-sdk
    ref: fc498f441485848f2c539ee66cd1f8983f842725
    remote: upstream
    tag: 1.3.2

docker:
  - from_image: veritone/aiware
    from_tag: test
    to_image: veritone/aiware-test
    to_tag: 1.6.2
    pull: no
```

Notice the tag value is overridden. When you run the tagger command:

```bash
tagger 0.1.0
```

Here is the explanation of commands:

### Git

- Change directory to `$GOPATH/src/github.com/veritone.com/veritone-sdk`
- `git checkout fc498f441485848f2c539ee66cd1f8983f842725`
- `git tag 1.3.2`
- `git push upstream 1.3.2`

### Docker

- Does not docker pull; assumes you have it locally
- `docker tag veritone/aiware:test veritone/aiware-test:1.6.2`
- `docker push veritone/aiware-test:1.6.2`

## Concurrency

Tagger offers a configurable concurrency using the `-c` flag. If you wanted to run 5 concurrent tagging tasks, here is the command:

```bash
tagger 0.1.0 -c 5
```

## Conclusion

In some usecases you need to tag everything the same version and you don't have mature pipelines. Tagger can easily satisfy this need for you. 😊

## About Veritone and Contact

Veritone is an AI platform where you can build, submit, and monetize AI engines for any business application. If you're interested in joining our team, see [Veritone Careers](https://www.veritone.com/about/careers/). 

If you have questions about this topic in particular, feel free to contact me via email: rsmith@veritone.com
