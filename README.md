# Go Experiments and Examples

Collection of assorted experiments and _starting points_ for Go.

In general, each directory has a self-contained experimental or starting point for a project.
Certain features, such as CI integrations, are done at the repo level (instead of directory) due to the requirements and
constraints from the different frameworks and integrations.

A _starting point_ project is the minimum set of source and configuration files for a project using some service, framework or library. Think of these as a combination of _"hello world"_ examples and templates. For example, suppose you want to build a (single) micro-service using gRPC. Then the starting project for this purpose contains the simplest definition and implementation of such service, something like the "echo service". It includes the corresponding gRPC/Probotobuf definitions,  client and server implementations, along with simple unit tests.

Notice that since the different dependencies and frameworks quickly evolve over time, thus the examples in this repo are
expected to be easily outdated.

## Contents

* Cannonical "Hello World" example in [hello](hello)

