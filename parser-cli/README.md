# Parser CLI

## Description
CLI tool to get word occurences in text.

## Installation
If you dont have a Golang setup, install go and setup the environment: 

Install [Go](https://go.dev/doc/install) as given in the official documentation.

Set GOPATH and update Path environment variable

```bash
export GOPATH=path/to/go/working/directory
export PATH=$PATH:$GOPATH/bin
```

Clone this project inside your GOPATH and install all project dependencies.

```bash
cd parser-cli
go mod download
```

## Usage
Make sure the api server is running. Please visit: https://github.com/pdhagat-git/golang-test-task/tree/main/app

Create a .env file in parser cli root folder and set the FILE_PATH and API_URL environment variables. Make sure FILE_PATH(optional) points to a .txt file and API_URL points to "/api/v1/parser/topWordCount" service

```bash
cp .env.example .env
```

Run cli app with filepath set in .env file.

```bash
go run main.go
```

Run cli app with optional path flag.

```bash
go run main.go --path path/to/file.txt
```

Optionally, you can also install the app to create a cli tool. 

```bash
go install
parser --path path/to/file.txt
```

## Licence
[![MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)