# Golang Test Task

## Description
App is an api service that has routes for getting top word occurences in text.
Parser-cli is a cli tool that lets you call the api service from command line. 

## Installation
If you dont have a Golang setup, install go and setup the environment: 

Install [Go](https://go.dev/doc/install) as given in the official documentation.

Set GOPATH and update Path environment variable.

```bash
export GOPATH=path/to/go/working/directory
export PATH=$PATH:$GOPATH/bin
```

Clone this project inside your GOPATH and install all project dependencies.

```bash
cd app
go mod download
```

```bash
cd parser-cli
go mod download
```
## Usage
Create a .env file in app root folder and set the PORT environment variable.

```bash
cp .env.example .env
```

Start the server.

```bash
go run main.go

```

Create a .env file in parser cli root folder and set the FILE_PATH and API_URL environment variables. Make sure FILE_PATH(optional) points to a .txt file and API_URL points to "/api/v1/parser/topWordCount" service.

```bash
cd parser-cli
cp .env.example .env
```

Install the app to create a cli tool. 

```bash
go install
parser --path path/to/file.txt
```

For API documentation read: https://github.com/pdhagat-git/golang-test-task/tree/main/app.
For CLI documentation read: 
## Licence
[![MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)