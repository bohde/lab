# lab

[![Build Status](https://travis-ci.org/joshbohde/lab.svg?branch=master)](https://travis-ci.org/joshbohde/lab)
[![GoDoc](https://godoc.org/github.com/joshbohde/lab?status.svg)](https://godoc.org/github.com/joshbohde/lab)

`lab` is a command line interface to [Gitlab](https://gitlab.com).

## Installation

```
$ go get -u github.com/joshbohde/lab
$ go install github.com/joshbohde/lab/cmd/lab
```

## Setup

1. Visit https://gitlab.com/profile/personal_access_tokens and create a new access token with API Scope
2. Run `git config --global lab.gitlab.com.token '<token>'`


## Commands

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
