package main

import "fmt"

func main() {
	//fmt.Printf("Hello %d!\n", "world"); // receives warning by linter
	fmt.Printf("Hello %s!\n", "world") // does not receive warning by linter
}
