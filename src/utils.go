package main

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
