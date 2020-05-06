Cachegreeter
============

Boring example web application that uses memcached.

This web application runs a hello world web service.

The index route "/" will display Hello World!

Using the "/name?name=" route you can change World to the name of your choice.

## Setup

Build dependencies
The go programming langauge must be installed and available at version 1.10 or newer.
With go installed, run `go get github.com/bradfitz/gomemcache` to install the only dependency outside of Go's standard library.

## Runtime dependencies

Memcached is required for this application to run. No other runtime dependencies are needed.

## Configuration

The web service defaults to running on port 8090. You can change the port by setting the PORT environment variable. For example to use port 8080 instead:
```
PORT=':8080' ./cachegreeter
```

The Memcache server location is configurable wth the MEMCACHE_URL environment variable. Memcache's default port on localhost is used if not specified.
For example:
```
MEMCACHE_URL="localhost:11211"


## Launch

go build

./cachegreeter
PORT=:7777 MEMCACHE_URL=localhost:11212 ./cachegreeter
