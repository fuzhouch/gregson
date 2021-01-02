# Gregson: A REST API Server Builder

This library provides a set of functions to create an web app, which
provides functionalities that can be accessed via a command line client
(behave like [awscli](https://github.com/aws/aws-cli)), or works with
an HTML5 Single Page Application.

Gregson does not aim to be a "framework" but a utility library to
use existing popular Go libraries, with a set of configured settings
to make web app reasonable to diagnose after it's deployed in
production.

Gregson is built on the popular Golang libraries like
[Gin](https://github.com/gin-gonic/gin),
[Zerolog](https://github.com/rs/zerolog), and
[Prometheus](https://github.com/prometheus/client_golang). Given
it's not framework, it directly expose data structures from the
libraries it references, with a purpose to avoid introduce new concepts
in code.

# Quick start

Gregson is built upon Go 1.14 or above, supporting ``go.mod`` mode only.
To reference Gregson, run commands below from your Go project.

```bash
    go get -v https://github.com/fuzhouch/gregson
```

A minimized example looks like below. It forces applying global logging
settings via ``InitGlobalZeroLog()``, then creates an ``gin.Engine``
object with ``zerolog`` and ``Prometheus`` hooked. Once ``g.Run()``
is called, developer can immediately see JSON log written in io.Stderr.
They can also use ``/metrics`` path to access Prometheus metrics.

```go

package main
import (
    "io"

    "github.com/gin-gonic/go"
    "github.com/fuzhouch/gregson"
)

func main() {
    gregson.InitGlobalZeroLog(io.Stderr)
    s := gregson.NewSetting()
    g := gregson.NewGin(s) // Zerolog and Prometheus are integrated.

    g.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "")
    })
    g.Run() // Listen and serve on 0.0.0.0:8080
}
```

Developer can use ``gregson.Setting`` structure to configure how
``gin.Engine`` object is created.

# About the name

The name of this library, "Gregson", refers to
[Tobias Gregson](https://www.arthur-conan-doyle.com/index.php/Tobias_Gregson),
the Scotland Yard inspector who worked with Sherlock Holmes in
[A Study of Scalet](https://www.arthur-conan-doyle.com/index.php?title=A_Study_in_Scarlet)
case.  He created his own theory to explain what happens, though
mostly wrong, which was named by Sherlock Holmes as "the smartest
detective in Scottland Yard".

## If you find it hard to customize... yes, it's by design

I pick name "Gregson" to reflect what we may find in this library: it
has its own attitude to select library and options which is needed in web
development scenario, even if this is not "correct" to everyone.
Customization is never part of design goals of
[Gregson](https://github.com/fuzhouch/gregson).

For example, [Gregson](https://github.com/fuzhouch/gregson) picks
[Gin](https://github.com/gin-gonic/gin) as web
framework and [zerolog](https://github.com/rs/zerolog) as logger.
If a developer wants to use [echo](https://github.com/labstack/echo)
and [logrus](https://github.com/sirupsen/logrus), don't submit a feature
request or PR. This library will not offer any customization to adopt
different frameworks. For same reason, if a developer want to change
default log level (which is forced to be ``zerolog.Info``), it will not
be accepted either.

The best option for these requests, is to copy useful snippets
from [Gregson](https://github.com/fuzhouch/gregson) to your code base.
It's completely allowed by MIT License.
