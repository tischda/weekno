# go-cli-template

Template for small [Go](https://www.golang.org) CLI projects.

## Get started

Name your project:
~~~
set PROJECT=my-project
~~~

Create repository and project folder with [Github CLI](https://github.com/cli/cli):
~~~
gh repo create %PROJECT% --confirm --public --template github.com/tischda/go-cli-template
cd %PROJECT%
go mod init github.com/tischda/%PROJECT%
go generate template.go
~~~

Start coding.

## Add modules

~~~
go mod tidy
go mod vendor
~~~

## Release project

~~~
make test
make release
~~~
