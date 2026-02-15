package main

import "fmt"

type Celsius float64
type Fahrenheit float64

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) String() string {
	return
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)

}

func FToC(f Fahrenheit) Celsius {
	return Celsius(((float64(f) - 32) * 5) / 9)
}

func main() {

}
