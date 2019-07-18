package tempconv

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// Exercise 2.1: Add types, constants, and functions to tempconv for processing
// temperatures in the Kelvin scale, where zero Kelvin is -273.15˚C and a
// difference of 1K has the same magnitude as 1˚C.

// KToC converts a Kelvin temperature to Celsius (exercise 2.1)
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
