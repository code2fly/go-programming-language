package tempconv

import "fmt"

// type to represent temp in celsius
type Celsius float64

// Fahrenheit type to represent temp in fahrenheit
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	freezingC     Celsius = 0
	boilingC      Celsius = 100
)

// convert temperature from Celcius to Farhenhite
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// convert temp from Farhenhite to Celsius
func FToC(f Fahrenheit) Celsius {
	// this is type convertion T(x)
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
