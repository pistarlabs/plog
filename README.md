# Plog

[![GoDoc](https://godoc.org/github.com/pistarlabs/plog?status.svg)](https://godoc.org/github.com/pistarlabs/plog) [![Build Status](https://travis-ci.org/pistarlabs/plog.svg?branch=master)](https://travis-ci.org/pistarlabs/plog) [![Build Status](https://drone.io/github.com/pistarlabs/plog/status.png)](https://drone.io/github.com/pistarlabs/plog/latest)

A simple HTTP request/response logger middleware for Go


## Usage
Using plog as logger middleware in [Julien Schmidt's httprouter](https://github.com/julienschmidt/httprouter)
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pistarlabs/plog"
)

func hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello world!")
}

func main() {
	router := httprouter.New()
	router.GET("/", hello)

	plog := plog.Default()

	http.ListenAndServe(":8080", plog.Handler(router))
}
```


Output in console
```
[log] 2015/09/22 21:55:49 |200|     49.662µs|[::1]:53679|GET /
```

See example folder for more example.

## Option
| Option 	| Type 	| Default 	| Description 	|
|:--------	|:--------	|:---------	|:----------------------------------	|
| Debug 	| bool 	| true 	| Output debug of log into console 	|
| Prefix 	| string 	| [log] 	| Prefix of log 	|

Example
```go
options := plog.Options{
  Debug:  true,
  Prefix: "[log ]",
}

plog := plog.New(options)
```
