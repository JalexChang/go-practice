# Concurrecny in GO

This is a folder storing my notes and source code of a book [Concurrency in Go: Tools and Techniques for Developers](https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197).

It sould be noticed that most of materials in this folder are extracted from the book and its source repo [concurrency-in-go-src](https://github.com/kat-co/concurrency-in-go-src) which is in MIT license.

## How to run sample code

Each sample code is a complete runnable source code that you can compile it and execute it independently and directlty.

Or, you can use the customized [Makefile](Makefile) to help you run the file easily by a command 
``` make run/test/bench NUM=<code_file_number> ```, which will automately find the corresponding source file and execute it via ```go run/test```

For example:

```
$ make run NUM=101
Find the code file in ./ch1-an-introduction-to-concurrency/101-race-condition.go
the value is 0.

$ make bench NUM=305
Find the code file in ./ch3-concurrency-building-blocks/goroutines/305-ctx-switch_test.go
goos: darwin
goarch: amd64
BenchmarkContextSwitch-4   	 5000000	       300 ns/op
PASS
ok  	command-line-arguments	1.813s

$ make test NUM=504
Find the code file in ./ch5-concurrency-at-scale/504-concurrent-with-hearbeat_test.go
=== RUN   TestDoWork_GeneratesAllNumbers
--- PASS: TestDoWork_GeneratesAllNumbers (2.00s)
PASS
ok  	command-line-arguments	2.009s
```