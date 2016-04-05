# Instant Whales

API Server to run/kill Docker containers instantly.
This server is for internal use. Don't publish the server on the Internet.

## Build

`make build`

## Run

`make run`

## Versions

```
$ go version    
go version go1.6 darwin/amd64

$ make --version
GNU Make 3.81

$ docker-compose version
docker-compose version 1.7.0rc1, build 1ad8866
docker-py version: 1.8.0-rc2

$ docker version
Client:
 Version:      1.9.0
 API version:  1.21
 Go version:   go1.4.3
 Git commit:   76d6bc9
 Built:        Tue Nov  3 19:20:09 UTC 2015
 OS/Arch:      darwin/amd64

Server:
 Version:      1.11.0-rc3
 API version:  1.23
 Go version:   go1.5.3
 Git commit:   eabf97a
 Built:        2016-04-01T23:33:49.977963402+00:00
 OS/Arch:      linux/amd64
```