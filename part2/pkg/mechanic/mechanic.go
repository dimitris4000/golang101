package mechanic

// Car abstracts the mechanical operations that can be performed on a car
type Car interface {
	OilOk() bool
	FillOil()
}

// ServiceCar will check a Car for mechanical issues and fixes them if required
func ServiceCar(car Car) {
	if !car.OilOk() {
		car.FillOil()
	}
}
