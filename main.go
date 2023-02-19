package main

import (
	"flag"
	"fmt"

	conv "github.com/onstrogithub/funtemps/conv"
)

var fahr float64
var celsius float64
var kelvin float64
var out string
var funfacts string

//var t string (funfacts)

func init() {
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader Celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader Kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	//flag.StringVar(&t, "t", "", "temperatur for fun-facts") (Eksempel til funfacts)
}

func main() {
	flag.Parse()

	if out == "C" && isFlagPassed("F") {
		celsius = conv.FarhenheitToCelsius(fahr)
		fmt.Printf("%.2f°F er %.2f°C\n", fahr, celsius)
	}
	if out == "F" && isFlagPassed("C") {
		fahr = conv.CelsiusToFarhenheit(celsius)
		fmt.Printf("%.2f°C er %.2f°F\n", celsius, fahr)
	}

	if out == "F" && isFlagPassed("K") {
		fahr = conv.KelvinToFarhenheit(kelvin)
		fmt.Printf("%.2f°K er %.2f°F\n", kelvin, fahr)
	}
	if out == "C" && isFlagPassed("K") {
		celsius = conv.KelvinToCelsius(kelvin)
		fmt.Printf("%.2f°K er %.2f°C\n", kelvin, celsius)
	}
	if out == "K" && isFlagPassed("C") {
		kelvin = conv.CelsiusToKelvin(celsius)
		fmt.Printf("%.2f°C er %.2f°K\n", celsius, kelvin)
	}
	if out == "K" && isFlagPassed("F") {
		kelvin = conv.FarhenheitToKelvin(fahr)
		fmt.Printf("%.2f°F er %.2f°K\n", fahr, kelvin)
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(k *flag.Flag) {
		if k.Name == name {
			found = true
		}
	})
	return found
}
