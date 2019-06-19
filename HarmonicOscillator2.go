package main

import (
	"fmt"
	"math"
	// "gonum.org/v1/plot"
	// "gonum.org/v1/plot/plotter"
	// "gonum.org/v1/plot/plotutil"
	// "gonum.org/v1/plot/vg"
)

type EquationInformation struct{
	InitialX, InitialY, InitialdY, InfinityX, InfinityY, InfinitydY, TargetX, TargetY, L float64
	SecondOrderDifferentalEquation func(x float64,y  float64, dy float64, p float64, l float64) float64
	InitialEquation func(x float64, l float64) float64
	InitialDifferentalEquation func(x float64, l float64) float64
	InfinityEquation func(x float64, p float64) float64
	InfinityDifferentalEquation func(x float64, p float64) float64
}

func main() {
	targetEquation := EquationInformation{
		InitialX: math.Pow(10, -5),
		InfinityX: 10,
		TargetX: 4,
		L: 1,
		InitialEquation: defineInitialEquation,
		InfinityEquation: defineInfinityEquation,
		InitialDifferentalEquation: defineInitialDifferntalEquation,
		InfinityDifferentalEquation: defineInfinityDifferentalEquation,
		SecondOrderDifferentalEquation: defineSecondOrderDifferentalEquation,
	}
	// examination
	targetEquation.SetInitialValues(1)
	fmt.Printf("%+v\n", targetEquation)
}

func defineSecondOrderDifferentalEquation(x float64, y float64, dy float64, p float64, l float64) float64 {
	return (l * (l + 1) * math.Pow(x, 2) + -2/x + math.Pow(p, 2))*y //l = 0, p = 1/n(n = 1)
}

func defineInitialEquation(x float64, l float64) float64 {
	return math.Pow(x, l+1)
}

func defineInitialDifferntalEquation(x float64, l float64) float64 {
	return 1 //l = 1
	//(l + 1) * math.Pow(x, l)
}

func defineInfinityEquation(x float64, p float64) float64 {
	return math.Pow(math.E, -p * x)
}

func defineInfinityDifferentalEquation(x float64, p float64) float64 {
	return -1 * p * math.Pow(math.E, -p * x)
}

func (equationInformation *EquationInformation) SetInitialValues(p float64) {
	equationInformation.InitialY = equationInformation.InitialEquation(equationInformation.InitialX, equationInformation.L)
	equationInformation.InitialdY = equationInformation.InitialDifferentalEquation(equationInformation.InitialX, equationInformation.L)
	equationInformation.InfinityY = equationInformation.InfinityEquation(equationInformation.InfinityX, p)
	equationInformation.InfinitydY = equationInformation.InfinityDifferentalEquation(equationInformation.InfinityX, p)
}
