## Why go is even required?

* infrastructure changed a lot where production code were being deployed
* multicore processors started becomming a more common thing
* cloud infrastructure with 100 / 1000 servers with multiprocessors became global
* infrastrcuture became
  * Scalable & distributed
  * Dynamic
  * More capacity

most langauge cannot make use of it more effeciently.

### Multithreading (Very Good feature/reason for Go)

* Concurrency without conflicts
* Complex code in java and expensive and slow
* Go was designed to run on multiple cores and built to support concurrency
* Concurrency in Go is cheap and easy

### Simple and Readable syntax of a dynamically typed language like python

### Efficiency and safety of a lower level statically typed language like c++

its mainly used at the server side or the backend side

Docker, Vault and kubernetes are wirtten in go.

* Simple Synatx , easy to read and write code
* Fast build time , start up and run
* Requires fewer resources
* compiled language
* consistent on multiple environments

  Go Programs must have one file with `package main and a main fucntion` to run.
  The "main" function is the entrypoint of a go program

  ##### All Go files start with package
* Defines the namespace for the code in the file, used to organize and group code
* main package is special , because it indicates that the file belongs to an executable program

### Slices in Go

example `var taskItems = []string{reviseGo, learnByCreating , reviseJava}`

Slices are dynamically-sized and more flexible than arrays

both slices and array can only hold elements of the same type
if we want to limit the items in the list then its going to be an array

`example : var taskItems = [3]string{reviseGo, learnByCreating , reviseJava} `

### Looping in GO

###### For Loop

* in Go lang we not only get the value of the item in the list but also the index of the item
* so we get two values in returns
