# Go Sandbox

For testing and learning in Go - starting 2019.06.03

> *Go CLI*

> `go build` --> Compiles a bunch of go source code files
> `go run` --> Compiles and executes one or two files
> `go fmt` --> Formats all the code in each file in the current directory
> `go install` --> Compiles nad "installs" a package
> `go get` --> Downloads the raw source code of someone else's package
> `go test` --> Runs any tests associated with the current project

---

*Packages*

`package main` defines a package that can be compiled and then executed - executable package
- must have a function called 'main'
For our purposes now, packages with altnernate names that can be used as dependencies but are not executable - reusable package

*Import statements*

Importing other packages - for example "fmt" is a package that's included in the go standard libary (https://golang.org/pkg/fmt/)

*Basic Go Types*
bool, string, int, float64

Arrays and Slices are similar dataytpes - every element in each must be the same type
arrays are fixed length, slices can grow and shrink. Zero-indexed.

