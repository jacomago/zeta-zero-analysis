package main

import (
	"fmt"
	"github.com/jacomago/zeta-zeros-analysis/data"
)

func main() {
	var f string
	fmt.Scan(&f)
	z := make(chan data.Zeros)
	go func() { z <- data.ImportZeros(f) }()

	fracDiff := make(chan []float64)

	go func() { fracDiff <- data.FractionalDifferences(<-z) }()

	d := <-fracDiff
	go func() {
		fmt.Println("The average is ", data.Avg(d))
	}()
	go func() {
		fmt.Println("The max is ", data.Max(d))
	}()
}
