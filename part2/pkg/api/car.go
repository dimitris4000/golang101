package api

import "fmt"

// Fiat is the programmatic representation of an actual car
type Fiat struct {
	oil        uint
	battery    uint
	engineRPMs uint
	Gas        uint
}

// NewFiat creates and initializes a new Fiat instance
func NewFiat(oil, battery, gas uint) *Fiat {
	return &Fiat{
		oil:     oil,
		battery: battery,
		Gas:     gas,
	}
}

// EngineStart will start the engine of the car if everything is ok with it
func (c *Fiat) EngineStart() {
	if c.engineRPMs > 0 {
		return
	}
	if c.oil <= 50 || c.battery < 20 || c.Gas < 10 {
		fmt.Println("Engine failed to start")
		return
	}
	c.Gas -= 2
	c.oil -= 10
	c.battery -= 10
	c.engineRPMs = 1000
	fmt.Println("Engine started")
}

// EngineStop will stop the engine of the car
func (c *Fiat) EngineStop() {
	fmt.Println("Stopping engine")
	c.engineRPMs = 0
}

// OilOk will check the oil level of the car and return true if all are good
func (c *Fiat) OilOk() bool {
	return c.oil > 50
}

// FillOil will set the Oil of the car to the correct level
func (c *Fiat) FillOil() {
	fmt.Println("Adding engine oil")
	c.oil = 100
}

// BatteryOk will check the battery of the car and return true if all are good
func (c *Fiat) BatteryOk() bool {
	return c.battery >= 20
}

// ChargeBattery will set the battery of the car to the correct charge level
func (c *Fiat) ChargeBattery() {
	fmt.Println("Charging car")
	c.battery = 100
}
