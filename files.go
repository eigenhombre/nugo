package main

import "fmt"

func Makefile(progname string) string {
	return fmt.Sprintf(`.PHONY: all install

all: %s

%s: main.go
	go build .

clean:
	rm -f %s

test:
	go test -v

install:
	go install .

docker:
	docker build -t %s .
`, progname, progname, progname, progname)
}

func main_go() string {
	return `package main

import (
	"fmt"
)

func main() {
	fmt.Println("OK")
}
`
}

func readme(progname string, year int) string {
	return fmt.Sprintf(`# %s

![build](https://github.com/eigenhombre/%s/actions/workflows/build.yml/badge.svg)

# License: MIT

Copyright (c) %d John Jacobsen

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

`, progname, progname, year)
}

func build() string {
	return `name: build

on:
  push:
  pull_request:
  workflow_dispatch:

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Build and test
        run: make docker
`
}

func Dockerfile(progname string) string {
	return fmt.Sprintf(`FROM golang:1.18

RUN apt-get -qq -y update
RUN apt-get -qq -y upgrade
RUN apt-get install -y make

WORKDIR /work

COPY . .

RUN make test %s
`, progname)
}

func gitignore(progname string) string {
	return fmt.Sprintf(`%s
`, progname)
}

func license(year int) string {
	return fmt.Sprintf(`MIT License

Copyright (c) %d John Jacobsen

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`, year)
}
