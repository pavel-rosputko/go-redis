package main

import "redis"
import "testing"
import __regexp__ "regexp"

var tests = []testing.InternalTest{
	{"redis.Test", redis.Test},
}
var benchmarks = []testing.InternalBenchmark{ //
	{"redis.BenchmarkPing", redis.BenchmarkPing},
	{"redis.BenchmarkSet", redis.BenchmarkSet},
}

func main() {
	testing.Main(__regexp__.MatchString, tests)
	testing.RunBenchmarks(__regexp__.MatchString, benchmarks)
}
