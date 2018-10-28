# Concurrecny in GO

This is a folder storing my notes and source code of a book [Concurrency in Go: Tools and Techniques for Developers](https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197).

It sould be noticed that most of materials in this folder are extracted from the book and its source repo [concurrency-in-go-src](https://github.com/kat-co/concurrency-in-go-src) which is in MIT license.

## How to run sample code

Each sample code in this folder is a complete runnable source code that you can compile it and execute it independently and directlty.

Or, you can use the customized [Makefile](Makefile) to help you run the file easily by a command 
``` make run NUM=<code_file_number> ```, which will automately find the corresponding source file, compile it, execute its binary, show a result, and then remove the binary.

For example:

```
$ make run NUM=101
Find the code file in /Users/jalexchang/Documents/coding-practice/concurrency-in-go/ch1-an-introduction-to-concurrency/101-race-condition.go
Execution Result:
---------

the value is 0.

---------
```