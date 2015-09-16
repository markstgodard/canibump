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

## Usage

    $ canibump
    Can I bump status is: YES

    $ echo $?
    0
