# nspectr
Query HTTP header and other DNS information.

# Installing

## Binaries
Binaries can be found [here](https://github.com/42wim/dt/releases/)

## Building
Go 1.6.3+ is required. Make sure you have [Go](https://golang.org/doc/install) properly installed, including setting up your [GOPATH] (https://golang.org/doc/code.html#GOPATH)

```
cd $GOPATH
go get github.com/geraldvillorente/nspectr
```

You should now have dt binary in the bin directory:

# Usage
```
Usage:
	nspectr -url=DOMAIN

Example:
	nspectr -cache=false -url=DOMAIN
	nspectr -url=DOMAIN age cache-control

Additional Parameters:
    age
        Return age only.
    cache-control
        Return cache-control only.
    x-cache
        Return x-cache only.
    set-cookie
        Return set-cookie only.
    strict-transport-security
        Return strict-transport-security only.
    etag
        Return etag only.
    x-served-by
        Return x-served-by only.

Flags:
  -cache
        To enable cache.
  -url
        Domain name to diagnose.
```