package main

import "testing"

// BenchmarkMain benchmarks the main function.
// NOTE: The go benchmark suite will run for at least a second. If the function
// finishes before a second, it will repeatdly run them until a seconds passes.
// Since there is a time.Sleep() call in the main function, this test will only
// run once as the main() function takes longer than a second. The flag
// `-benchtime=10s` can be given to increase the minimum time to benchmark
// (increasing the runs per benchmark) to get more consistent results between
// benchmarks.
func BenchmarkMain(b *testing.B) {
	// run the main function b.N times
	for n := 0; n < b.N; n++ {
		main()
	}
}

// TestMain tests the main function
// NOTE: This function itself doesn't test for anything but it can be used to
// add conditions that main() (or any other function) should meet.
func TestMain(t *testing.T) {
	// Doesn't test for anything. Will always pass
	main()
}
