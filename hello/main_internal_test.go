package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello, world!
}

func TestGreet(t *testing.T) {
	want := "Hello, world!"

	got := greet()

	if got != want {
		t.Errorf("Expected: %q, Got: %q", want, got)
	}
}
