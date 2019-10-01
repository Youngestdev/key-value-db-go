package main

import (
	"fmt"
	keyValueDB "key-value-db-go/src"
)

// Default or should I say starter methods.

func main() {
	fmt.Println("Key value stuff going on here.")
	test := keyValueDB.ValueObject{}
	test.SetKey("Laisi2")
	test.SetVal("123")
	test.Test(0)
	fmt.Println(test.GetKey())
	fmt.Println(test.GetH())
	fmt.Println(test.ToString())
	fmt.Println(test.IsChar(","))
}
