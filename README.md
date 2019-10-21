
![Course title image](https://raw.githubusercontent.com/JoshuaMorris/programming-is-for-lovers/master/img/p4L_banner.jpg)

# Programming for lovers

This project contains my notes and homework assignments from the Programming is for lovers course

### Running tests with code coverage
to run the test suite, execute `go test -race -count=1 -cover ./...` from the root of the project.
- pass the `update` flag to rewrite the golden files

```console
🍟 🌄 : go test -race -count=1 -cover ./...
?   	github.com/JoshuaMorris/programming-for-lovers/lectures/lecture-1-go-basics/conditionals	[no test files]
ok  	github.com/JoshuaMorris/programming-for-lovers/lectures/lecture-1-go-basics/gcd	18.438s	coverage: 100.0% of statements
?   	github.com/JoshuaMorris/programming-for-lovers/lectures/lecture-1-go-basics/intro	[no test files]
?   	github.com/JoshuaMorris/programming-for-lovers/lectures/lecture-1-go-basics/loops	[no test files]
ok  	github.com/JoshuaMorris/programming-for-lovers/lectures/lecture-1-go-basics/primes	1.109s	coverage: 100.0% of statements
ok  	github.com/JoshuaMorris/programming-for-lovers/testutils	1.073s	coverage: 100.0% of statements

```
 

### running function benchmarks
To run the function benchmarks, execute `go test ./... -bench=.` from the root of the project

```console
🍏 🌄 : go test ./... -bench=.
?   	github.com/JoshuaMorris/programming-for-lovers/lectures/lecture-1-go-basics/conditionals	[no test files]
goos: darwin
goarch: amd64
pkg: github.com/JoshuaMorris/programming-for-lovers/lectures/lecture-1-go-basics/gcd
BenchmarkTrivialGCD-12        	   13198	    119493 ns/op
BenchmarkAnotherSumEven-12    	  198627	    124092 ns/op
BenchmarkGCD-12               	32691813	        36.6 ns/op
BenchmarkEuclidGCD-12         	  195224	    114009 ns/op
PASS
ok  	github.com/JoshuaMorris/programming-for-lovers/lectures/lecture-1-go-basics/gcd	53.466s
?   	github.com/JoshuaMorris/programming-for-lovers/lectures/lecture-1-go-basics/intro	[no test files]
?   	github.com/JoshuaMorris/programming-for-lovers/lectures/lecture-1-go-basics/loops	[no test files]
goos: darwin
goarch: amd64
pkg: github.com/JoshuaMorris/programming-for-lovers/lectures/lecture-1-go-basics/primes
BenchmarkTrivialPrimeFinder-12     	   10000	   1229966 ns/op
BenchmarkIsPrime-12                	100000000	        12.1 ns/op
BenchmarkSieveOfEratosthenes-12    	   75721	    125987 ns/op
BenchmarkCrossOffMultiples-12      	236972158	         4.40 ns/op
PASS
ok  	github.com/JoshuaMorris/programming-for-lovers/lectures/lecture-1-go-basics/primes	24.884s
PASS
ok  	github.com/JoshuaMorris/programming-for-lovers/testutils	0.067s

```
