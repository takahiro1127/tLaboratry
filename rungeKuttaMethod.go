package main

import (
	"fmt"
)

type EquationInformation struct{ //中身はコマンド叩くときに持ってこれるとなお良い
	InitialX, InitialY float64
	DifferentalEquation func(x float64,y  float64) float64
}
func main() {
	targetEquation := EquationInformation{
		InitialX: 0,
		InitialY: 1,
		DifferentalEquation: defineEquationInformation,
	}
	fmt.Println(targetEquation.rungeKuttaMethod(1))
}

func defineEquationInformation(x float64, y float64) float64 { //TODO コマンド実行時に式を指定できるようにする
	return y
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


