# Basic Course to start a Go App with Docker.

Firstable clone the repository from github.

## To Run directly on a Docker container exec this:

```
# make docker-build
# make docker-run
```

To see output, check container logs.
For more commads check the makefile and /docs.


## To Run locally (needs a GoLang installation):

If you want to run this project without docker:
1. Install latest Golang version from here.
2. Execute this on terminal

```
# go mod init
# go mod tidy
```

```
# make run
```

