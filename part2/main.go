package main

import (
	"fmt"
	"golang101/pkg/api"
)

func main() {
	myCar := api.NewFiat(40, 10, 100)
	fmt.Println(myCar)

	fmt.Println()
	myCar.EngineStart()
	fmt.Println(myCar)
}
