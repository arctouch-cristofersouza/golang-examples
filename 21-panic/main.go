package main

import "os"

// panic throw a fatal error which stops everything if not treated and
// a common use of panic is to abort if a function returns an error value that we don’t know how to (or want to) handle
func main() {

	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
