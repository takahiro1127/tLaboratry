package main

import (
	"fmt"
	"math"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type EquationInformation struct{ //中身はコマンド叩くときに持ってこれるとなお良い
	InitialX, InitialY, InitialdY float64
	SecondOrderDifferentalEquation func(x float64,y  float64, dy float64, p float64) float64
}
func main() {
	targetEquation := EquationInformation{
		InitialX: 0,
		InitialY: 1,
		InitialdY: 1,
		SecondOrderDifferentalEquation: defineSecondOrderDifferentalEquation,
	}
	fmt.Println(targetEquation.SecondOrderRungeKuttaMethod(1, 1.0))
}

func defineSecondOrderDifferentalEquation(x float64, y float64, dy float64, p float64) float64 { //TODO コマンド実行時に式を指定できるようにする
	return (-2/x + math.Pow(p, 2))*y //l = 0, p = 1/n(n = 1)
}

func (equationInformation *EquationInformation) SecondOrderRungeKuttaMethod(targetX int, p float64) float64{
	var targetY = equationInformation.InitialY
	var targetdY = equationInformation.InitialdY
	stockYvalue := [][]float64{}
	for i := int(equationInformation.InitialX*1000); i < targetX*1000; i++ {
		var step = 0.001

		var p1 = step*equationInformation.SecondOrderDifferentalEquation(equationInformation.InitialX, targetY, targetdY, p)
		var k1 = step*targetdY

		var p2 = step*equationInformation.SecondOrderDifferentalEquation(equationInformation.InitialX+step/2, targetY+k1/2, targetdY+p1/2, p)
		var k2 = step*(targetdY+p1/2)

		var p3 = step*equationInformation.SecondOrderDifferentalEquation(equationInformation.InitialX+step/2, targetY+k2/2, targetdY+p2/2, p)
		var k3 = step*(targetdY+p2/2)

		var p4 = step*equationInformation.SecondOrderDifferentalEquation(equationInformation.InitialX+step, targetY+k3, targetdY+p3, p)
		var k4 = step*(targetdY+p3)

		equationInformation.InitialX = equationInformation.InitialX + step
		targetY = targetY + (k1 + 2 * k2 + 2 * k3 + k4)/6
		targetdY = targetdY + (p1 + 2 * p2 + 2 * p3 + p4)/6
		if i%33 == 0 {
			fmt.Println(stockYvalue)
			stockYvalue = append(stockYvalue, []float64{float64(i/1000)+float64(i%1000)/1000, targetY})
		}
	}
	err := makeGraphFromSecondOrderArray(stockYvalue, "fromZero")
	if !err {
	panic(err)
	}
	return targetY
}

func makeGraphFromSecondOrderArray(arr [][]float64, title string) bool {
	p, err := plot.New()
	if err != nil {
			panic(err)
	}
	p.Title.Text = "hydrogenAtom:" + title
	p.X.Label.Text = "r axis"
	p.Y.Label.Text = "φ axis"

	pts := make(plotter.XYs, len(arr))

	for i, axis := range arr {
			pts[i].X = axis[0]
			pts[i].Y = axis[1]
	}
	err = plotutil.AddLinePoints(p, pts)
	if err != nil {
			panic(err)
	}
	if err := p.Save(5*vg.Inch, 5*vg.Inch, title + ".png"); err != nil{
			panic(err)
	}
	return err == nil
}

