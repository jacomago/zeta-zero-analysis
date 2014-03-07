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

func Max(values []float64) float64 {
	result := 0.0
	for _, v := range values {
		result = math.Max(v, result)
	}
	return result
}

func Avg(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
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

func ExportData(data []float64) {
}

func floatsEqual(a float64, b float64) bool {
	Epsilon := 0.0000001
	if math.Abs(a-b) < Epsilon {
		return true
	}
	return false
}

// fractional differences
func closestZero(z Zeros, zero float64) float64 {
	diff := make(chan float64)
	minDistance := 1.0
	go func() {
		for _, decimalZero := range z {
			if !floatsEqual(decimalZero, zero) {
				diff <- math.Abs(decimalZero - zero)
			}
		}
		close(diff)
	}()
	for d := range diff {
		minDistance = math.Min(d, minDistance)
	}
	return minDistance
}

func FractionalDifferences(z Zeros) []float64 {
	distance := make(chan float64)
	result := make([]float64, 0, len(z))
	go func() {
		for _, zero := range z {
			distance <- closestZero(z, zero)
		}
		close(distance)
	}()

	for d := range distance {
		result = append(result, d)
	}
	return result
}

// zeros epsilon close to other
func fCloseTo(z Zeros, zero float64, epsilon float64) int {
	result := 0
	test := make(chan bool)
	go func() {
		for _, decimalZero := range z {
			test <- math.Abs(decimalZero-zero) < epsilon
		}
		close(test)
	}()
	for b := range test {
		if b {
			result += 1
		}
	}
	return result
}

type Psi func(int) float64

func ZerosFCloseTo(z Zeros, p Psi) []int {
	value := make(chan int)
	values := make([]int, 0, len(z))
	go func() {
		for _, zero := range z {
			value <- fCloseTo(z, zero, p(len(z)))
		}
		close(value)
	}()
	for v := range value {
		values = append(values, v)
	}
	return values
}

// histogram creating
func createHistogram(z Zeros, epsilon float64) map[int]float64 {
	return nil
}
