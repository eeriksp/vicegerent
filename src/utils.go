package main

// A convenience function to make panicking more succinct.
func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
