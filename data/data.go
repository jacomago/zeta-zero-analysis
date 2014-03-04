package data

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

type Zeros map[float64]float64

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ImportZeros(filename string) Zeros {
	f, err := os.Open(filename)
	check(err)

	z := make(Zeros)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		number, parseErr := strconv.ParseFloat(text, 64)
		z[number] = number - math.Floor(number)
		check(parseErr)
	}
	check(scanner.Err())
	return z
}

func ExportData(data) {
}

func floatsEqual(a float64, b float64) bool {
	Epsilon := 0.000001
	if math.Abs(a-b) < Epsilon {
		return true
	}
	return false
}

// fractional differences
func closestZero(z Zeros, zero float64) float64 {
}

func fractionalDifferences(z Zeros) (float64, float64) {
}

// zeros epsilon close to other
func fCloseTo(z Zeros, zero float64, epsilon float64) int {
}

func zerosFCloseTo(z Zeros, epsilon float64) float64 {
}

// histogram creating
func createHistogram(z Zeros, epsilon float64) map[int]float64 {
}
