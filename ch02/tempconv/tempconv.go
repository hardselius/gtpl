package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g˚C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g˚F", f)
}

// Exercise 2.1: Add types, constants, and functions to tempconv for processing
// temperatures in the Kelvin scale, where zero Kelvin is -273.15˚C and a
// difference of 1K has the same magnitude as 1˚C.

type Kelvin float64

const (
	AbsoluteZeroK Kelvin = 0
	FreezingK     Kelvin = 273.15
	BilingK       Kelvin = 373.15
)

func (k Kelvin) String() string {
	return fmt.Sprintf("%g˚K", k)
}
