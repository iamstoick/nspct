# nspct
Query HTTP headers and other DNS information.

# Installing

## Binaries
Coming soon...

## Building
Go 1.6.3+ is required. Make sure you have [Go](https://golang.org/doc/install) properly installed, including setting up your [GOPATH] (https://golang.org/doc/code.html#GOPATH)

```
cd $GOPATH
go get github.com/geraldvillorente/nspct
```

## Make the command globally available
You should now have `nspct` binary in the bin directory: `$GOPATH/bin`. 
To make `$GOPATH/bin` globally available so you can execute the command `nspct`anywhere in the terminal/console:
* Open your ~/.bash_profile file.
* Add the `$GOPATH/bin` path at the bottom like this `export PATH=$PATH:$GOPATH/bin`. 
* Save and close.
* Reload `~/.bash_profile` by running `. ~/.bash_profile` or `source ~/.bash_profile`in the terminal/console.

# Usage
```
Usage:
	nspct -url=DOMAIN

Example:
	nspct -url=DOMAIN -cache
	nspct -url=DOMAIN age cache-control

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
        To disanable cache bypass.
  -url
        Domain name to diagnose.
```