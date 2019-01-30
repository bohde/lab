# lab

[![Build Status](https://travis-ci.org/joshbohde/lab.svg?branch=master)](https://travis-ci.org/joshbohde/lab)
[![GoDoc](https://godoc.org/github.com/joshbohde/lab?status.svg)](https://godoc.org/github.com/joshbohde/lab)

`lab` is a command line interface to [Gitlab](https://gitlab.com).

## Installation

### Release

1. Download [the latest release](https://github.com/joshbohde/lab/releases/latest) for your platform.
2. Make it executable, e.g. `chmod +x lab_linux_amd64`
3. Put it somewhere on your path, e.g. `mv lab_linux_amd64 /usr/local/bin/lab`

### From source

```
$ go get -u github.com/joshbohde/lab
$ go install github.com/joshbohde/lab/cmd/lab
```

## Commands

### `auth`

Will configure access tokens for the current project, if none exist.

```
$ lab auth
```

This will open your browser to your Gitlab access tokens page, for either https://gitlab.com, or your self-hosted instance. Create a new token with API scope, and paste it back into your terminal.

### `issue`

Opens an issue in the current project.

```
$ lab issue
```

This will open your editor so that you can fill in the title and description of the issue.

Flags are available for command line scripting. To see a full list of available flags, run `lab issue --help`.

### `merge-request`

Opens a merge request from the current branch to the default branch.

```
$ lab merge-request
```

This will open your editor so that you can fill in the title and description of the merge request.

Flags are available for command line scripting. To see a full list of available flags, run `lab merge-request --help`.
