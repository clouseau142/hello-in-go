package main

import "fmt"

func main() {
	name := "Frank-2"
	fmt.Printf("hello, world\n")
	fmt.Printf(name + "\n")
	fmt.Printf(Meinprint() + "\n")
}

// returns mein
func Meinprint() string {
	return "mein"
}
