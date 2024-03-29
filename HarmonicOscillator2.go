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
	InitialDifferentalEquation func(x float64, l float64) float64
	InfinityEquation func(x float64, p float64) float64
	InfinityDifferentalEquation func(x float64, p float64) float64
}
type Wronskian struct {
	targetY, targetdY, targetYfromInfinity, targetdYfromInfinity float64
}

type Wronskians []Wronskian

func main() {
	targetEquation := EquationInformation{
		InitialX: math.Pow(10, -5),
		InfinityX: 20,
		TargetX: 2,
		L: 0,
		InitialEquation: defineInitialEquation,
		InfinityEquation: defineInfinityEquation,
		InitialDifferentalEquation: defineInitialDifferntalEquation,
		InfinityDifferentalEquation: defineInfinityDifferentalEquation,
		SecondOrderDifferentalEquation: defineSecondOrderDifferentalEquation,
	}

	var p = []float64{0.51, 0.5}
	var wronskians Wronskians
	for i := 0; i < 30; i++ {
		targetEquation.SetInitialValues(p[i])
		wronskian := Wronskian{}
		wronskian.targetY, wronskian.targetdY = targetEquation.SecondOrderRungeKuttaMethod(p[i], true, false)
		wronskian.targetYfromInfinity, wronskian.targetdYfromInfinity = targetEquation.SecondOrderRungeKuttaMethod(p[i], false, false)
		wronskians = append(wronskians, wronskian)
		if i == 0 {
			continue
		}
		beforeWronskian := wronskians[i - 1].Determinant()
		currentWronskian := wronskian.Determinant()
		if i == 4 {
			fmt.Println(wronskian)
			fmt.Println(wronskians[i])
		}
		p = append(p, ((currentWronskian * p[i -1] - beforeWronskian * p[i]) / (currentWronskian - beforeWronskian)))
		if (math.Abs((p[i] - p[i - 1])/p[i -1])) < math.Pow(10, -5) {
			targetEquation.SecondOrderRungeKuttaMethod(p[i], true, true)
			targetEquation.SecondOrderRungeKuttaMethod(p[i], false, true)
			break
		}
	}
	fmt.Println(p)
}

func defineSecondOrderDifferentalEquation(x float64, y float64, dy float64, p float64, l float64) float64 {
	return (l * (l + 1) * math.Pow(x, 2) + -2/x + math.Pow(p, 2))*y //l = 0, p = 1/n(n = 1)
}

func defineInitialEquation(x float64, l float64) float64 {
	return math.Pow(x, l+1)
}

func defineInitialDifferntalEquation(x float64, l float64) float64 {
	// return (l + 1) * math.Pow(x, l)
	return 1
}

func defineInfinityEquation(x float64, p float64) float64 {
	return math.Pow(math.E, -1 * p * x)
	// return math.Pow(math.E, -1 * p * x)
}

func defineInfinityDifferentalEquation(x float64, p float64) float64 {
	return -1 * p * math.Pow(math.E, -1 * p * x)
}

func (equationInformation *EquationInformation) SetInitialValues(p float64) {
	equationInformation.InitialY = equationInformation.InitialEquation(equationInformation.InitialX, equationInformation.L)
	equationInformation.InitialdY = equationInformation.InitialDifferentalEquation(equationInformation.InitialX, equationInformation.L)
	equationInformation.InfinityY = equationInformation.InfinityEquation(equationInformation.InfinityX, p)
	equationInformation.InfinitydY = equationInformation.InfinityDifferentalEquation(equationInformation.InfinityX, p)
}

func (wronskian *Wronskian) Determinant() float64 {
	return wronskian.targetY * wronskian.targetdYfromInfinity - wronskian.targetdY * wronskian.targetYfromInfinity
}

func (equationInformation *EquationInformation) SecondOrderRungeKuttaMethod(p float64, plus bool, makeGraph bool) (float64, float64) {
	var targetY = equationInformation.InitialY
	var targetX = equationInformation.TargetX
	var targetdY = equationInformation.InitialdY
	var initialX = equationInformation.InitialX
	var stockYvalue = [][]float64{{equationInformation.InitialX, equationInformation.InitialY}}

	if !plus {
		targetY = equationInformation.InfinityY
		targetX = equationInformation.TargetX
		targetdY = equationInformation.InfinityY
		initialX = equationInformation.InfinityX
		stockYvalue = [][]float64{{equationInformation.InfinityX, equationInformation.InfinityY}}
	}

	for i := int(initialX*1000); i != int(targetX*1000); i++ {
		var step = 0.001
		var p1, k1, p2, k2, p3, k3, p4, k4 float64
		if plus {
			p1 = step*equationInformation.SecondOrderDifferentalEquation(initialX, targetY, targetdY, p, equationInformation.L)
			k1 = step*targetdY
			p2 = step*equationInformation.SecondOrderDifferentalEquation(initialX+step/2, targetY+k1/2, targetdY+p1/2, p, equationInformation.L)
			k2 = step*(targetdY+p1/2)
			p3 = step*equationInformation.SecondOrderDifferentalEquation(initialX+step/2, targetY+k2/2, targetdY+p2/2, p, equationInformation.L)
			k3 = step*(targetdY+p2/2)
			p4 = step*equationInformation.SecondOrderDifferentalEquation(initialX+step, targetY+k3, targetdY+p3, p, equationInformation.L)
			k4 = step*(targetdY+p3)
			initialX = initialX + step
			targetY = targetY + (k1 + 2 * k2 + 2 * k3 + k4)/6
			targetdY = targetdY + (p1 + 2 * p2 + 2 * p3 + p4)/6
		} else {
			p1 = step*equationInformation.SecondOrderDifferentalEquation(initialX, targetY, targetdY, p, equationInformation.L)
			k1 = step*targetdY
			p2 = step*equationInformation.SecondOrderDifferentalEquation(initialX-step/2, targetY+k1/2, targetdY+p1/2, p, equationInformation.L)
			k2 = step*(targetdY-p1/2)
			p3 = step*equationInformation.SecondOrderDifferentalEquation(initialX-step/2, targetY+k2/2, targetdY+p2/2, p, equationInformation.L)
			k3 = step*(targetdY-p2/2)
			p4 = step*equationInformation.SecondOrderDifferentalEquation(initialX-step, targetY+k3, targetdY+p3, p, equationInformation.L)
			k4 = step*(targetdY-p3)
			initialX = initialX - step
			targetY = targetY + (k1 + 2 * k2 + 2 * k3 + k4)/6
			targetdY = targetdY + (p1 + 2 * p2 + 2 * p3 + p4)/6
			i = i - 2
		}
		if i%33 == 0 {
			stockYvalue = append(stockYvalue, []float64{float64(i/1000)+float64(i%1000)/1000, targetY})
		}
	}
	if makeGraph {
		var title = "fromZero"
		if !plus { title = "fromInfinity" }
		err := makeGraphFromSecondOrderArray(stockYvalue, title)
		if !err {
			panic(err)
		}
		fmt.Println(p)
		fmt.Println("作ったよ")
	}
	if !plus {
		return targetY, -1 * targetdY
	}
	return targetY, targetdY
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

