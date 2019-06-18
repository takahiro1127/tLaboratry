package main

import (
	"fmt"
	"math"
)

type hydrogenAtomEquationInformation struct{
	MaximumX, yMaximumX, MinimumX, yMinimumX, float64
	DifferentalEquation func(x float64,y  float64) float64
	Equation func(x float64) float64
}

func main() {
	targetEquation := EquationInformation{
		MaximumX: 5,
		MinimumX: -5,
		DifferentalEquation: defineDifferentalEquationInformation,
		Equation: defineEquationInformation,
	}
	targetEquation.rungeKuttaMethod(1)
}

func defineEquationInformation(x float64) float64 { //TODO コマンド実行時に式を指定できるようにする
	return math.Pow(math.E, -1*math.Pow(x, 2)/2)
}

func defineDifferentalEquationInformation(x float64, y float64) float64 { //TODO コマンド実行時に式を指定できるようにする
	return -1*x*math.Pow(math.E, -1*math.Pow(x, 2)/2)
}

func (equationInformation *EquationInformation)setYvalueFromEquation {}

func (equationInformation *EquationInformation) rungeKuttaMethod() float64{
}

func (equationInformation *EquationInformation) rungeKuttaMethod(targetX int) float64{
	var targetY = equationInformation.InitialY
	for i := int(equationInformation.InitialX*1000); i < targetX*1000; i++ {
		var step = 0.001
		var k1 = step*equationInformation.DifferentalEquation(equationInformation.InitialX, targetY)
		var k2 = step*equationInformation.DifferentalEquation(equationInformation.InitialX + step/2, targetY + k1/2)
		var k3 = step*equationInformation.DifferentalEquation(equationInformation.InitialX + step/2, targetY + k2/2)
		var k4 = step*equationInformation.DifferentalEquation(equationInformation.InitialX + step, targetY + k3)
		equationInformation.InitialX = equationInformation.InitialX + step
		targetY = targetY + (k1 + 2 * k2 + 2 * k3 + k4)/6
	}
	return targetY
}
