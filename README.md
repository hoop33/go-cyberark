# go-cyberark

> A Go wrapper for the CyberArk Vault API

[![Build Status](https://travis-ci.org/hoop33/go-cyberark.svg?branch=master)](https://travis-ci.org/hoop33/go-cyberark)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](http://opensource.org/licenses/MIT)
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/hoop33/go-cyberark)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoop33/go-cyberark)](https://goreportcard.com/report/github.com/hoop33/go-cyberark)

## Table of Contents

* [Introduction](#introduction)
* [Installation](#installation)
* [Usage](#usage)
* [Contributing](#contributing)
* [Credits](#credits)
* [Disclaimer](#disclaimer)
* [License](#license)

## Introduction

go-cyberark is a client library to talk to the CyberArk Vault API. Note that this library is purpose-built for my specific use case, and may not cover your use case. Contributions welcome -- see **Contributing** below.

## Installation

You must have a working Go installation. To install, type:

```sh
$ go get -u github.com/hoop33/go-cyberark
```

You then import the library with this import path:

```go
import cyberark "github.com/hoop33/cyberark"
```

## Usage

To use go-cyberark, create a client, get one of its services, and call its `Do()` function.

Note that the only service currently offered is `GetPassword`.

Example:

```go
client, err := cyberark.NewClient(
  cyberark.SetHost("cyberark.example.com"),
)
if err != nil {
  log.Fatal(err.Error())
}

ret, err := client.GetPassword().
  AppID("my_app_id").
  Safe("my_safe").
  Object("LDAP").
  Do()
if err != nil {
  log.Fatal(err.Error())
}

if ret.ErrorCode != "" {
  log.Fatal(ret.ErrorCode)
}

log.Println(ret.UserName)
log.Println(ret.Content)
```

Look in the `examples` folder for examples you can run.

## Contributing

Please note that this project is released with a [Contributor Code of Conduct](http://contributor-covenant.org/). By participating in this project you agree to abide by its terms. See [CODE_OF_CONDUCT](CODE_OF_CONDUCT.md) file.

Contributions are welcome! Please open pull requests with code that passes all the checks. See *Building* for more information.

## Building

You must have a working Go development environment. The included `Makefile` is used for building. To get started, run:

```sh
$ make deps
```

This will install necessary dependencies. You should have to do this only once, or when you upgrade Go.

To run tests and build, run:

```sh
$ make
```

To get coverage reports, run:

```sh
$ make coverage
```

This will open a browser with the coverage reports.

## Credits

* [Testify](https://github.com/stretchr/testify)

Also, a hat tip to [Oliver Elihard](https://github.com/olivere) and his [elastic](https://github.com/olivere/elastic) project. I borrowed heavily from this project's approach, architecture, and code. Any faults are my own.

## Disclaimer

go-cyberark is not affiliated with, nor endorsed by, CyberArk Software Ltd.

## License

Copyright &copy; 2017 Rob Warner

Licensed under the [MIT License](https://hoop33.mit-license.org/)

