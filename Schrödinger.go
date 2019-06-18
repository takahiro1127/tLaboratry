package main

import (
	"fmt"
	"math"
)

func main() {
	var energy int
	fmt.Println("Energy:")
	fmt.Scan(&energy)
	var ramda int
	ramda = calcRamdaFromEnergy(energy)
	targetEquation := EquationInformation{
		InitialX: 0,
		InitialY: 1,
		InitialdY: 1,
		SecondOrderDifferentalEquation: defineEquation,
	}
}

func calcRamdaFromEnergy(energy int) {
	return 
}

func defineEquation(x float64, y float64, dy float64) float64 { //TODO コマンド実行時に式を指定できるようにする
	return y
}

type EquationInformation struct{ //中身はコマンド叩くときに持ってこれるとなお良い
	InitialX, InitialY, InitialdY float64
	SecondOrderDifferentalEquation func(x float64,y  float64, dy float64) float64
}

func (equationInformation *EquationInformation)  DimenssionlessSchrödingerEquation(targetX int) float64{
	equationInformation.InitialY = (math.E)^-5

}

func (equationInformation *EquationInformation) SecondOrderRungeKuttaMethod(targetX int) float64{
	var targetY = equationInformation.InitialY
	var targetdY = equationInformation.InitialdY
	for i := int(equationInformation.InitialX*1000); i < targetX*1000; i++ {
		var step = 0.001

		var p1 = step*equationInformation.SecondOrderDifferentalEquation(equationInformation.InitialX, targetY, targetdY)
		var k1 = step*targetdY

		var p2 = step*equationInformation.SecondOrderDifferentalEquation(equationInformation.InitialX+step/2, targetY+k1/2, targetdY+p1/2)
		var k2 = step*(targetdY+p1/2)

		var p3 = step*equationInformation.SecondOrderDifferentalEquation(equationInformation.InitialX+step/2, targetY+k2/2, targetdY+p2/2)
		var k3 = step*(targetdY+p2/2)

		var p4 = step*equationInformation.SecondOrderDifferentalEquation(equationInformation.InitialX+step, targetY+k3, targetdY+p3)
		var k4 = step*(targetdY+p3)

		equationInformation.InitialX = equationInformation.InitialX + step
		targetY = targetY + (k1 + 2 * k2 + 2 * k3 + k4)/6
		targetdY = targetdY + (p1 + 2 * p2 + 2 * p3 + p4)/6
	}
	return targetY
}
