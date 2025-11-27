# WUTL

WULT(web utilities) Is simple library with some nice to have helpers when working with [GO](https://go.dev) http servers made with the standard package. This library may be compatible with other frameworks but I would not be sure since I don't use them. I have only work with the standard http package.


## Adding middle ware to a `http.Handler`

```go

import (
  "net/http"
  "github.com/chapgx/wutl"
)

// create a handler that you will eventually pass to your http server 

mx :=  http.NewServerMux()
handler := wutl.NewHandler(mx)

// you can pass as many Handler functions as you want they are executed left to right or top to bottom.
handler.AddMiddleware(log1, statistics)
```


## Nice to have middleware

Serving embedded files. You can define a root director that will be the base access, anything outside root will be treated as private. You can pass as many 
skip functions, which are custom rules in case you want to skip serving from an embedded file system for specific resources.

```go
handler := wutl.NewHandler(mux)

handler.AddMiddleware(wutl.ServeEmbedded(embeddedFS, "root_dir", skipfunc))

```



Serving static files. The standard library has a function that does the same. I wrote my own in case I want to change the rules inside of it later. As of now it just serves static files. Same logic as embedded for root directory define, nothing outside of of the specify root is allowed access.

```go
handler := wutl.NewHandler(mux)

handler.AddMiddleware(wutl.ServerFiles("rood_dir"))
```


