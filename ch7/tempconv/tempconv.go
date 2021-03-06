package tempconv

import (
	"flag"
	"fmt"

	"github.com/yykhomenko/book-gopl/ch2/tempconv"
)

type celsiusFlag struct{ tempconv.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var value float64
	var unit string
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°С":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	default:
		return fmt.Errorf("incorrect temperature: %s", s)
	}
}

func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
