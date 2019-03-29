# Go 1.9 Examples
A collection of examples written in the Go language

## Useful resources

https://tour.golang.org/

https://gobyexample.com/

https://golangbot.com/

https://golang.org/doc/ is the official documentation, and also provides links to direct source code of all standard Go functions or any that have been incorporated from third parties.


## Grains of Salt

These examples were provided by a student, not a Go master.  Further, they were written to bridge the gap between how it looks in Java vs Go.  They are not necessarily good representations of how to perform these actions in the Go language.

Further, this was written when Go 1.9 was the newest version.  A lot of these examples compare to the Java version, and Java is big on reflection and generics.  Go 1.9 notably lacks these in the same robustness, and the examples reflect that.


## Completed

Legend:

* X - Done
* N - N/A (Refer to README in that directory)
* ? - Possibly inapplicable to Go

- [X] Lambdas
- [ ] Misc
- [ ] Multicasting
- [ ] Naming
- [ ] Reflection
- [ ] RMI/RPC
  - [X] ex1-HelloServer
  - [X] ex2-SquareServer
  - [X] ex3-Client-Callback
  - [X] ex4-Asynchronous-Server
  - [ ] ex5-Load-Remote-Class
  - [ ] ex6-RMI-Thread-Safety
  - [ ] ex7-PassingArgsInRMI
  - [ ] ex8-with-timeout
  - [ ] ex9-HelloServer-2-interfaces
- [ ] Security
- [ ] Serialization
  - [ ] TODO: Go does not support non-identical inner-struct conversions
    - [X] MioAlma
	- [X] MioAlmaDos
- [ ] Sockets
  - [ ] LargerHTTPd
  - [X] TCP
     - [X] Client
     - [X] MultiThreaded
     - [X] ObjectServer
     - [X] SingleThreaded
     - [N] SocketOptions
  - [ ] TinyHTTPd
  - [X] UDP
- [X] Threads
  - [X] TestAccount
  - [N] SynchronizedAccount (Implemented in TestAccount/Account)
  	- [X] Account
  - [X] Consumer
  - [X] Producer
  - [N] Remaining Examples (Should be relatively self-evident)
