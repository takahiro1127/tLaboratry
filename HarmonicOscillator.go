package main

import (
	"fmt"
	"math"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

//散乱状態を求めるアルゴリズムを組んで、実際の実験データと比較する感じ
//多次元空間においてるんげ食った
//共役勾配法


type EquationInformation struct{ //中身はコマンド叩くときに持ってこれるとなお良い
	InitialX, InitialY, InitialdY, InfinityX, InfinityY, InfinitydY, TargetX, TargetY, L float64
	SecondOrderDifferentalEquation func(x float64,y  float64, dy float64, p float64, l float64) float64
	InitialEquation func(x float64, l float64) float64
	InfinityEquation func(x float64, p float64) float64
}

var p = []float64{0.8, 0.9}
var wronskian = [][][]float64{}

func main() {
	targetEquation := EquationInformation{
		InitialX: math.Pow(10, -5),
		InitialdY: 1,
		InfinityX: 10,
		TargetX: 4,
		L: 1,
		InfinitydY: 100 * math.Pow(math.E, -10),
		InitialEquation: defineInitialEquation,
		InfinityEquation: defineInfinityEquation,
		SecondOrderDifferentalEquation: defineSecondOrderDifferentalEquation,
	}

	for i := 0; i < 30; i++ {
		targetEquation.SetInitialYvalue()
		targetEquation.SetInfinityYvalue(p[i])
		var targetY, targetdY = targetEquation.SecondOrderRungeKuttaMethod(targetEquation.TargetX, p[i], false)
		var targetYfromInfinity, targetdYfromInfinity = targetEquation.SecondOrderRungeKuttaMethodFromInfinity(targetEquation.TargetX, p[i], false)
		wronskian = append(wronskian, [][]float64{{targetY, targetdY}, {targetYfromInfinity, targetdYfromInfinity}})
		// fmt.Println(wronskian)
		if i == 0 { continue  }

		wronskianIminus1 := wronskian[i-1][0][0]*wronskian[i-1][1][1] - wronskian[i-1][0][1]*wronskian[i-1][1][0]
		if i == 3 {
			fmt.Println(wronskian)
		}
		wronskianI := wronskian[i][0][0]*wronskian[i][1][1] - wronskian[i][0][1]*wronskian[i][1][0]
		if i == 3 {
			fmt.Println(wronskianIminus1)
		}

		if i == 30 {
			fmt.Println(wronskian)
		}
		p = append(p, Round((wronskianI * p[i-1] - wronskianIminus1 * p[i])/(wronskianI - wronskianIminus1), 4))


		if Round(targetY, 2) == Round(targetYfromInfinity, 2) && Round(targetdY, 2) == Round(targetdYfromInfinity, 2) {
			targetEquation.SecondOrderRungeKuttaMethod(targetEquation.TargetX, p[i], true)
			targetEquation.SecondOrderRungeKuttaMethodFromInfinity(targetEquation.TargetX, p[i], true)
			break
		}
	}
	// fmt.Println(targetEquation.SecondOrderRungeKuttaMethod(targetEquation.TargetX, 1, false))
}

func Round(val float64, place int) float64 {
	shift := math.Pow(10, float64(place))
 	return math.Floor(val * shift + .5) / shift
}

func defineSecondOrderDifferentalEquation(x float64, y float64, dy float64, p float64, l float64) float64 { //TODO コマンド実行時に式を指定できるようにする
	return (l * (l + 1) * math.Pow(x, 2) + -2/x + math.Pow(p, 2))*y //l = 0, p = 1/n(n = 1)
}

func defineInitialEquation(x float64, l float64) float64 { //TODO コマンド実行時に式を指定できるようにする
	return math.Pow(x, l+1) //r**l+1 (l=0)
}

func defineInfinityEquation(x float64, p float64) float64 { //TODO コマンド実行時に式を指定できるようにする
	return p * math.Pow(math.E, -p * x) //r**l+1 (l=0)
}

func (equationInformation *EquationInformation) SetInitialYvalue() {
	equationInformation.InitialY = equationInformation.InitialEquation(equationInformation.InitialX, equationInformation.L)
}

func (equationInformation *EquationInformation) SetInfinityYvalue(p float64) {
	equationInformation.InfinityY = equationInformation.InfinityEquation(equationInformation.InfinityX, p)
}

func (equationInformation *EquationInformation) SecondOrderRungeKuttaMethodFromInfinity(targetX float64, p float64, makeGraph bool) (float64, float64) {
	var targetY = equationInformation.InfinityY
	var targetdY = equationInformation.InfinitydY
	stockYvalue := [][]float64{}


	for i := int(equationInformation.InfinityX*1000); i > int(targetX*1000); i-- {
		var step = 0.001

		var p1 = step*equationInformation.SecondOrderDifferentalEquation(equationInformation.InitialX, targetY, targetdY, p, equationInformation.L)
		var k1 = step*targetdY

		var p2 = step*equationInformation.SecondOrderDifferentalEquation(equationInformation.InitialX+step/2, targetY+k1/2, targetdY+p1/2, p, equationInformation.L)
		var k2 = step*(targetdY+p1/2)

		var p3 = step*equationInformation.SecondOrderDifferentalEquation(equationInformation.InitialX+step/2, targetY+k2/2, targetdY+p2/2, p, equationInformation.L)
		var k3 = step*(targetdY+p2/2)

		var p4 = step*equationInformation.SecondOrderDifferentalEquation(equationInformation.InitialX+step, targetY+k3, targetdY+p3, p, equationInformation.L)
		var k4 = step*(targetdY+p3)

		equationInformation.InitialX = equationInformation.InitialX + step
		targetY = targetY + (k1 + 2 * k2 + 2 * k3 + k4)/6
		targetdY = targetdY + (p1 + 2 * p2 + 2 * p3 + p4)/6
		if i%33 == 0 {
			stockYvalue = append(stockYvalue, []float64{float64(i/1000)+float64(i%1000)/1000, targetY})
		}
	}
	if makeGraph {
		err := makeGraphFromSecondOrderArray(stockYvalue, "fromInfinity")
		if !err {
			fmt.Println()
			panic(err)
		}
	}
	return targetY, targetdY
}



func (equationInformation *EquationInformation) SecondOrderRungeKuttaMethod(targetX float64, p float64, makeGraph bool) (float64, float64) {
	var targetY = equationInformation.InitialY
	var targetdY = equationInformation.InitialdY
	stockYvalue := [][]float64{}


	for i := int(equationInformation.InitialX*1000); i < int(targetX*1000); i++ {
		var step = 0.001

		var p1 = step*equationInformation.SecondOrderDifferentalEquation(equationInformation.InitialX, targetY, targetdY, p, equationInformation.L)
		var k1 = step*targetdY

		var p2 = step*equationInformation.SecondOrderDifferentalEquation(equationInformation.InitialX+step/2, targetY+k1/2, targetdY+p1/2, p, equationInformation.L)
		var k2 = step*(targetdY+p1/2)

		var p3 = step*equationInformation.SecondOrderDifferentalEquation(equationInformation.InitialX+step/2, targetY+k2/2, targetdY+p2/2, p, equationInformation.L)
		var k3 = step*(targetdY+p2/2)

		var p4 = step*equationInformation.SecondOrderDifferentalEquation(equationInformation.InitialX+step, targetY+k3, targetdY+p3, p, equationInformation.L)
		var k4 = step*(targetdY+p3)

		equationInformation.InitialX = equationInformation.InitialX + step
		targetY = targetY + (k1 + 2 * k2 + 2 * k3 + k4)/6
		targetdY = targetdY + (p1 + 2 * p2 + 2 * p3 + p4)/6
		if i%33 == 0 {
			stockYvalue = append(stockYvalue, []float64{float64(i/1000)+float64(i%1000)/1000, targetY})
		}
	}
	if makeGraph {
		err := makeGraphFromSecondOrderArray(stockYvalue, "fromInfinity")
		if !err {
			fmt.Println()
			panic(err)
		}
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
