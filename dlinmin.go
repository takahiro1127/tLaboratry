package main

import(
	"fmt"
)

type Vector struct {

}

// Aには[[行要素m個]*n個]でn*m行列を作る


func main() {
	
}

func conjugarteGradient(A [][]int, b []int, x []int) []int{
	dimension := len(b)
	r := make([]int, dimension)
	p := make([]int, dimension)
	a := make([]int, dimension)
	r[0] = minus(b, product(A, x))
	p[0] = r[0]
	for i := 0; i < dimension; i++ {
		a[i] = 
	}
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

func conjugate(x, y []int, A [][]int) {
	
}


// りょうし多体問題のシュレディンガー方程式をとく 平均場の理論ハードリーフォーク　密度汎関数理論 多体問題を一体問題に帰着させる
//
