# hasuraenv

CLI to manage multiple hasura-cli versions.

## Installation

```sh
$ go install github.com/kmtym1998/hasuraenv/cmd/hasuraenv@latest
$ hasuraenv init
$ export PATH=~/.hasuraenv/current:$PATH;
```

## Usage

```shell
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
