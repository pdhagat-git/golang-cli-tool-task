# Parser Service

## Description
API server to parse text files and return word occurences.

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
cd app
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

## REST API
Install curl to make http requests over command line.

```bash
sudo apt install curl 
```

Get occurences of all words in text.

```bash
curl -d '{"data": "Your text"}' -H 'Content-Type:application/json' http://localhost:PORT/api/v1/parser/wordCount
```

Get top 10 word occurences in text.

```bash
curl -d '{"data": "Your text"}' -H 'Content-Type:application/json' http://localhost:PORT/api/v1/parser/topWordCount
```

## Licence
[![MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)