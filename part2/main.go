package main

import (
	"fmt"
	"golang101/pkg/api"
	"golang101/pkg/electric"
)

func main() {
	myCar := api.NewFiat(40, 10, 100)
	fmt.Println(myCar)

	fmt.Println()
	myCar.EngineStart()
	fmt.Println(myCar)

	fmt.Println()
	electric.ServiceCar(myCar)
	fmt.Println(myCar)
	myCar.EngineStart()
	fmt.Println(myCar)
}
