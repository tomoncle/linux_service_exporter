# linux_service_exporter
[![travis-ci.org](https://travis-ci.org/tomoncle/linux_service_exporter.svg?branch=master)](https://travis-ci.org/tomoncle/linux_service_exporter) [![Go Report Card](https://goreportcard.com/badge/github.com/tomoncle/linux_service_exporter)](https://goreportcard.com/report/github.com/tomoncle/linux_service_exporter) ![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/tomoncle/linux_service_exporter.svg) ![GitHub repo size](https://img.shields.io/github/repo-size/tomoncle/linux_service_exporter.svg?color=green&logoColor=green) ![GitHub top language](https://img.shields.io/github/languages/top/tomoncle/linux_service_exporter.svg?color=yes) ![GitHub license](https://img.shields.io/github/license/tomoncle/linux_service_exporter.svg)

Prometheus linux service exporter.

### require
* go version : `1.17.0+`
* prometheus : `2.0.0+`
* os         : `linux`

### build
```bash
$ git clone https://github.com/tomoncle/linux_service_exporter
$ cd linux_service_exporter && go mod download
$ go build linux_service_exporter.go
```
