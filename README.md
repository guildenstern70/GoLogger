# GOLOGGER PROJECT

[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)

Go Logger is a fake, fancy logger, that logs absurd sentences.

### Setup

You should have a single workspace repository for all of your GO
code, and a GOPATH pointing to that.

It is advisable to have a /src directory inside your workspace.

This code must be installed in a directory inside

    $GOPATH/src

with the following path

    github.com/guildenstern70/golearn

so that you find the LICENSE file here:

    $GOPATH/src/github.com/guildenstern70/gologger/LICENSE

### Build

You build the executable file by running

    go install github.com/guildenstern70/gologger/cmd

### Run

Run the built assembly with

    $GOPATH/bin/cmd.exe (Windows)
    $GOPATH/bin/cmd (Mac + Linux)

### Test
Run the test suite of this project by running

    go test github.com/guildenstern70/gologger/core