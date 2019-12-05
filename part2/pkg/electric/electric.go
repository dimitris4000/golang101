package electric

// Car abstracts the electric operations that can be performed on a car
type Car interface {
	BatteryOk() bool
	ChargeBattery()
}

// ServiceCar will check a Car for electric issues an fix if required
func ServiceCar(car Car) {
	if !car.BatteryOk() {
		car.ChargeBattery()
	}
}
