# Gregson: A web development utility library with attitude

A collection of tools that helps creating a standardized web application.
It is not a "framework", but a set of tools to integrate existing
libraries with a set of pre-defined default configuration, which gives
a consistent behavior when building a web application.

## Dependencies

Gregson is built on the following libraries:

- [Gin](https://github.com/gin-gonic/gin)
- [Zerolog](https://github.com/rs/zerolog)

# About this name

The name of this library, "Gregson", refers to
[Tobias Gregson](https://www.arthur-conan-doyle.com/index.php/Tobias_Gregson),
the Scotland Yard inspector who worked with Sherlock Holmes in
[A Study of Scalet](https://www.arthur-conan-doyle.com/index.php?title=A_Study_in_Scarlet)
case.  He created his own theory to explain what happens, though
mostly wrong, which was named by Sherlock Holmes as "the smartest
detective in Scottland Yard".

# It appears hard to customize... yes, it's by design

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
Don't worry about legal issue. It's completely allowed via MIT License.
