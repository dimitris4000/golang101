package main

import "fmt"

func main() {
	map1 := map[string]bool{
		"key1": true,
		"key2": true,
	}

	map1["key3"] = false // create and set a new key:value pair

	fmt.Println(map1)

	delete(map1, "key2")

	fmt.Println(map1)
}
