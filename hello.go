package main

import "fmt"

func main() {
	name := "Frank"
	fmt.Printf("Hallo, Welt\n")
	fmt.Printf(name + "\n")
	fmt.Printf(Meinprint() + "\n")
}

// Meinprint returns a string
func Meinprint() string {
	return "mein auf deutsch gewechselt"
}
