# hasuraenv

[![Latest release](https://img.shields.io/github/v/release/kmtym1998/hasuraenv)](https://github.com/kmtym1998/hasuraenv/releases/latest)
[![Test](https://github.com/kmtym1998/hasuraenv/actions/workflows/test.yaml/badge.svg)](https://github.com/kmtym1998/hasuraenv/actions/workflows/test.yaml)
[![Static Check](https://github.com/kmtym1998/hasuraenv/actions/workflows/static-check.yaml/badge.svg)](https://github.com/kmtym1998/hasuraenv/actions/workflows/static-check.yaml)
[![Vulnerability Check](https://github.com/kmtym1998/hasuraenv/actions/workflows/govulncheck.yaml/badge.svg)]
[![Go Report Card](https://goreportcard.com/badge/github.com/kmtym1998/hasuraenv)](https://goreportcard.com/report/github.com/kmtym1998/hasuraenv)
![GitHub license](https://img.shields.io/github/license/kmtym1998/hasuraenv)

Hasura CLI version manager

## Installation

```
$ go install github.com/kmtym1998/hasuraenv/cmd/hasuraenv@latest
$ hasuraenv init
$ export PATH=~/.hasuraenv/current:$PATH;
```

## Usage

```
$ hasuraenv ls-remote --limit 10
INFO Latest 10 releases
     v2.13.0-beta.1
     v2.12.0
     ...

$ hasuraenv install v2.13.0-beta.1

$ hasuraenv ls
INFO Installed hasura cli
     v2.1.0
     v2.9.0
     v2.9.0-beta.3
     v2.13.0-beta.1

$ hasuraenv use v2.13.0-beta.1

$ hasura version
INFO hasura cli                                    version=v2.13.0-beta.1
```
