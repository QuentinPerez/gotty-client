# gotty-client
:wrench: [GoTTY](https://github.com/yudai/gotty) client for your terminal

![](https://raw.githubusercontent.com/moul/gotty-client/master/resources/gotty-client.png)

[![Build Status](https://travis-ci.org/moul/gotty-client.svg?branch=master)](https://travis-ci.org/moul/gotty-client)
[![GoDoc](https://godoc.org/github.com/moul/gotty-client?status.svg)](https://godoc.org/github.com/moul/gotty-client)

```ruby
                +----------------+       +----------------+      +-------------+
                |                |       |             +--------->  /bin/bash  |
            +--->    Browser    -----+   |     gotty   |  |      +-------------+
+-------+   |   |                |   |   |             |  |
|       |   |   +----------------+   |   |             |  |      +-------------+
|  Bob  +---+                        +--->---websockets+--------->  /bin/bash  |
|       |   |   +================+   |   |             |  |      +-------------+
+-------+   |   |................|   |   |             |  |
            +--->..gotty-client.-----+   |             |  |      +-------------+
                |................|       |             +--------->  /bin/bash  |
                +================+       +----------------+      +-------------+

                  ^  ^  ^  ^  ^
                  |  |  |  |  |
```

## Example

Server side ([GoTTY](https://github.com/yudai/gotty))

```console
$ gotty -p 9191 sh -c 'while true; do date; sleep 1; done'
2015/08/24 18:54:31 Server is starting with command: sh -c while true; do date; sleep 1; done
2015/08/24 18:54:31 URL: http://[::1]:9191/
2015/08/24 18:54:34 GET /ws
2015/08/24 18:54:34 New client connected: 127.0.0.1:61811
2015/08/24 18:54:34 Command is running for client 127.0.0.1:61811 with PID 64834
2015/08/24 18:54:39 Command exited for: 127.0.0.1:61811
2015/08/24 18:54:39 Connection closed: 127.0.0.1:61811
...
```

**Client side**

```console
$ gotty-client http://localhost:9191/
INFO[0000] New title: GoTTY - sh -c while true; do date; sleep 1; done (jean-michel-van-damme.local)
WARN[0000] Unhandled protocol message: json pref: 2{}
Mon Aug 24 18:54:34 CEST 2015
Mon Aug 24 18:54:35 CEST 2015
Mon Aug 24 18:54:36 CEST 2015
Mon Aug 24 18:54:37 CEST 2015
Mon Aug 24 18:54:38 CEST 2015
^C
```

## Usage

```console
$ gotty-client -h
NAME:
   gotty-client - GoTTY client for your terminal

USAGE:
   gotty-client [global options] command [command options] GOTTY_URL

VERSION:
   1.1.0

AUTHOR(S):
   Manfred Touron <https://github.com/moul/gotty-client>

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

## Install

Install latest version using Golang (recommended)

```console
$ go get github.com/moul/gotty-client/cmd/gotty-client
```

---

Install latest version using Homebrew (Mac OS X)

```console
$ brew install https://raw.githubusercontent.com/moul/gotty-client/master/contrib/homebrew/gotty-client.rb --HEAD
```

or the latest released version

```console
$ brew install https://raw.githubusercontent.com/moul/ssh2docker/master/contrib/homebrew/assh.rb
```

## Changelog

### master (unreleased)

* Add debug mode (`--debug`/`-D`) ([#18](https://github.com/moul/gotty-client/issues/18))
* Improve error message when connecting by checking HTTP status code
* Fix arguments passing ([#16](https://github.com/moul/gotty-client/issues/16))

[full commits list](https://github.com/moul/gotty-client/compare/v1.1.0...master)

### [v1.1.0](https://github.com/moul/gotty-client/releases/tag/v1.1.0) (2015-10-10)

* Handling arguments + using mutexes (thanks to [@QuentinPerez](https://github.com/QuentinPerez))
* Add logo ([#9](https://github.com/moul/gotty-client/issues/9))
* Using codegansta/cli for CLI parsing ([#3](https://github.com/moul/gotty-client/issues/3))
* Fix panic when running on older GoTTY server ([#13](https://github.com/moul/gotty-client/issues/13))
* Add 'homebrew support' ([#1](https://github.com/moul/gotty-client/issues/1))
* Add Changelog ([#5](https://github.com/moul/gotty-client/issues/5))
* Add GOXC configuration to build binaries for multiple architectures ([#2](https://github.com/moul/gotty-client/issues/2))

[full commits list](https://github.com/moul/gotty-client/compare/v1.0.1...v1.1.0)

### [v1.0.1](https://github.com/moul/gotty-client/releases/tag/v1.0.1) (2015-09-27)

* Using party to manage dependencies

[full commits list](https://github.com/moul/gotty-client/compare/v1.0.0...v1.0.1)

### [v1.0.0](https://github.com/moul/gotty-client/releases/tag/v1.0.0) (2015-09-27)

Compatible with [GoTTY](https://github.com/yudai/gotty) version: [v0.0.10](https://github.com/yudai/gotty/releases/tag/v0.0.10)

#### Features

* Support **basic-auth**
* Support **terminal-(re)size**
* Support **write**
* Support **title**
* Support **custom URI**

[full commits list](https://github.com/moul/gotty-client/compare/cf0c1146c7ce20fe0bd65764c13253bc575cd43a...v1.0.0)

## License

MIT
