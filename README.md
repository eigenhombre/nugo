# nugo

![build](https://github.com/eigenhombre/nugo/actions/workflows/build.yml/badge.svg)

This small program serves as my Go project generator.  If you use it, you'll have to
change the GitHub username to your own.

Aside from a simple Go 'hello world' type program, it creates some
infrastructure I tend to use for all my builds: Dockerfile, Makefile,
GitHub Action workflow, etc.

Here's how I build it:

    make
    make install

Here's how I test it:

    make test

(This *deletes* and replaces any existing `nugotestproject` on $GOPATH.)

Example run:

    $ nugo myproject
    go: creating new go.mod: module github.com/eigenhombre/myproject
    go: to add module requirements and sums:
        go mod tidy
    Initialized empty Git repository in /Users/jacobsen/Programming/go/myproject/.git/


	       go.mod      LICENSE   Dockerfile     Makefile
	    README.md   .gitignore      .github         .git
	      main.go

    OK
    $ cd $GOPATH/myproject
    $ make
    go build .
    $ make install
    go install .
    $ cd
    $ myproject
    OK
    $

# License: MIT

Copyright (c) 2023 John Jacobsen

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
