# vanity

[![Go Report Card](https://goreportcard.com/badge/github.com/clly/vanity)](https://goreportcard.com/report/github.com/clly/vanity)
[![Build Status](https://travis-ci.org/clly/vanity.svg?branch=travis)](https://travis-ci.org/clly/vanity)
[![Godoc](https://godoc.org/github.com/clly/vanity?status.svg)](https://godoc.org/github.com/clly/vanity)
[![License](https://img.shields.io/github/license/clly/vanity.svg)](LICENSE)

Vanity is a CLI program to create vanity buttons in a github readme for Go 
programs. It also creates .travis.yml files.

## Installation

```
go get -u github.com/clly/vanity
go install github.com/clly/vanity
```

OR

```
go get github.com/clly/vanity
go get -u github.com/golang/dep/cmd/dep
cd $GOPATH/github.com/clly/vanity
dep ensure
go install
```

## Usage

We're assuming that this is a github project and that the package matches the github name.
```
vanity github.com/clly/vanity
```
Output
```
# vanity

[![Go Report Card](https://goreportcard.com/badge/github.com/clly/vanity)](https://goreportcard.com/report/github.com/clly/vanity)
[![Build Status](https://travis-ci.org/clly/vanity.svg?branch=travis)](https://travis-ci.org/clly/vanity)
[![Godoc](https://godoc.org/github.com/clly/vanity?status.svg)](https://godoc.org/github.com/clly/vanity)
[![License](https://img.shields.io/github/license/clly/vanity.svg)](LICENSE)
```
