// Write a function that returns the last rune of a string.

package main

import "github.com/01-edu/z01"

func LastRune(s string) rune {
	runs := []rune(s)
	return runs[len(runs)-1]
}

func main() {
	z01.PrintRune(LastRune("HelloZ"))
	z01.PrintRune(10)
}
