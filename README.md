# About
This repository serves as an example of how profiling code can be added to a
golang application so that it can generate profiling output if the correct
flags are passed.

## Build
After cloning, simply type `go build` to build the application.

## Execute With Profiling
Run the application with profiling flags. e.g:
`./go_profile_example -cpuprofile cpu.prof -memprofile mem.prof`

This will run the application and generate two profiling files `cpu.prof` and
`mem.prof`.

## View Profile
[`pprof`](https://golang.org/pkg/runtime/pprof/) is a profiling tool that is
quite useful for viewing the generated profiles. To view profiles in a web
browser one can execute the following in the build directory:

`pprof -http localhost:5000 go_profile_example cpu.prof`

This will start a server on the local machine at port 5000. A web page should
then automatically open to display the profile graph. For more output options
see the documentation of [`pprof`](https://golang.org/pkg/runtime/pprof/).

The golang blog on [profiling go programs](https://blog.golang.org/profiling-go-programs) 
also provides a nice explanation on how to profile Go applications.

Last Modified Tue 19 Feb 09:36:33 GMT 2019
