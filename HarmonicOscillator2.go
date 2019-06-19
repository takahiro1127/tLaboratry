package main

import (
	"fmt"
	"math"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type EquationInformation struct{
	InitialX, InitialY, InitialdY, InfinityX, InfinityY, InfinitydY, TargetX, TargetY, L float64
	SecondOrderDifferentalEquation func(x float64,y  float64, dy float64, p float64, l float64) float64
	InitialEquation func(x float64, l float64) float64
	InfinityEquation func(x float64, p float64) float64
}

func main() {
	targetEquation := EquationInformation{
		InitialX: math.Pow(10, -5),
		InfinityX: 10,
		TargetX: 4,
		L: 1,
		InitialEquation: defineInitialEquation,
		InfinityEquation: defineInfinityEquation,
		SecondOrderDifferentalEquation: defineSecondOrderDifferentalEquation,
	}
}

func defineSecondOrderDifferentalEquation(x float64, y float64, dy float64, p float64, l float64) float64 {
	return (l * (l + 1) * math.Pow(x, 2) + -2/x + math.Pow(p, 2))*y //l = 0, p = 1/n(n = 1)
}

func defineInitialEquation(x float64, l float64) float64 {
	return math.Pow(x, l+1)
}

func defineInitialDifferntalEquation(x float64, l float64) float64 {
	return (l + 1) * math.Pow(x, l)
}

func defineInfinityEquation(x float64, p float64) float64 {
	return math.Pow(math.E, -p * x)
}

func defineInfinityDifferentalEquation(x float64, p float64) float64 {
	return -1 * p * math.Pow(math.E, -p * x)
}

func (equationInformation *EquationInformation) SetInitialYvalue() {
	equationInformation.InitialY = equationInformation.InitialEquation(equationInformation.InitialX, equationInformation.L)
}

func (equationInfotmation *EquationInformation) SetInitialdYValue() {
	equationInformation.InitialY = equationInformation.InitialDifferntalEquation(equationInformation.InitialX, equationInformation.L)
}

func (equationInformation *EquationInformation) SetInfinityYvalue(p float64) {
	equationInformation.InfinityY = equationInformation.InfinityEquation(equationInformation.InfinityX, p)
}

func (equationInformation *EquationInformation) SetInfinityYvalue(p float64) {
	equationInformation.InfinityY = equationInformation.InfinityDiffrentalEquation(equationInformation.InfinityX, p)
}
