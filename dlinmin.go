package main

import(
	"fmt"
)

// Aには[[行要素m個]*n個]でn*m行列を作る


func main() {
	var A [][]int
	A[0] = []int{3, 0}
	A[1] = []int{0, 4}
	x := []int{3, 9}
	b := []int{1, 45}
	fmt.Println(conjugarteGradient(A, b, x))
}

func conjugarteGradient(A [][]int, b []int, x []int) []int{
	dimension := len(b)
	r := make([]int, dimension)
	p := make([]int, dimension)
	a := make([]int, dimension)
	B := make([][]int, dimension)
	targetx := make([][]int, dimension) //自信ないところ！！！
	targetx[0] = x
	B[0] = b
	r[0] = minus(b, product(A, x))
	p[0] = r[0]
	for i := 0; i < dimension; i++ {
		a[i] = product(r[i], p[i])/(product(p[i], product(A, p[i])))
		x[i+1] = x[i] + a[i]*r[i]
		r[i+1] = r[i] - a[i]*product(A, p[i])
		B[i] = -1 * product(r[i+1], product(A, p[i]))/(product(p[i], product(A, p[i])))
		p[i+1] = r[i+1] + b[i]*p[k]
	}
	return x[len(x)]
}

func product(A [][]int, x []int) []int{
	dimension := len(x)
	if len(A[0]) != dimension { panic("matrix dimenstion not match") }
	ax := make([]int, dimension)
	for i := 0; i < len(A); i++ {
		for j := 0; j < dimension; j++ {
			ax[j] = A[i][j] * x[j]
		}
	}
}

func minus(a, b []int) []int{
	dimension = len(a)
	if len(b) != dimension { panic("matrix dimenstion not match") }
	x := make([]int, dimension)
	for i := 0; i < dimension; i++ {
		x[i] = a[i] - b[i]
	}
	return x
}

func product(x, y []int) []int {
	dimension = len(x)
	xy := make([]int, dimension)
	for j := 0; j < dimension; j++ {
		xy[j] = y[j] * x[j]
	}
	return xy
}


// りょうし多体問題のシュレディンガー方程式をとく 平均場の理論ハードリーフォーク　密度汎関数理論 多体問題を一体問題に帰着させる
//
