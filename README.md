[![Build Status](https://travis-ci.org/markstgodard/canibump.svg?branch=master)](https://travis-ci.org/markstgodard/canibump)
# CAN I BUMP ?

## About
canibump is a simple command-line tool to check if
[canibump website](https://canibump.cfapps.io) is **YES** or **NO**.

*Note: You can use this in a CI/CD pipeline by checking the exit code:
- 0 - if you can bump OR
- 1 - if cannot bump



## Install

    go get github.com/markstgodard/canibump

## Usage: command line

    $ canibump
    Can I bump status is: YES

    $ echo $?
    0

## Usage: CI/CD Concourse
Example using canibump in a [Concourse](http://concourse.ci) CI/CD pipeline

**pipeline.yml**
````yaml
- name: bump-it
  serial: true
  plan:
  - aggregate:
    - get: cf-release-src
      passed: [run-acceptance-tests]
    - get: ci-scripts
  - task: can-i-bump
    privileged: true
    file: ci-scripts/routing-ci/scripts/ci/check-can-i-bump.yml
  - task: bump
    privileged: true
    file: ci-scripts/routing-ci/scripts/ci/bump.yml
  - put: cf-release-src
````

**check-can-i-bump.yml**
````yaml
---
platform: linux

image: docker:///cloudfoundry/runtime-ci

inputs:
  - name: ci-scripts

run:
  path: ci-scripts/routing-ci/scripts/ci/check-can-i-bump

````

**check-can-i-bump**
````sh
#!/bin/bash
set -e -x

export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH

export GOPATH=$PWD
export PATH=$GOPATH/bin:$PATH

go get github.com/markstgodard/canibump

````
