package main

import (
	"errors"
	"fmt"
)

func Divide(a, b float64) (result float64, err error) {
	if b == 0 {
		return 0.0, errors.New("The divisor can't be 0")
	}
	result = a / b
	return result, nil
}

func SafeDivideAll(nums []float64, divisor float64) ([]float64, []error) {
	results := []float64{}
	errs := []error{}

	for _, value := range nums {
		result, err := Divide(value, divisor)

		results = append(results, result)
		errs = append(errs, err)
	}

	return results, errs
}

func printAll(results []float64, err []error) {
	for i := range results {
		fmt.Printf("Result: %v, Error: %v \n", results[i], err[i])
	}
}

func main() {
	nums := []float64{1, 10, 2, 9, 7.5, 4.6, 8.45}
	divisor1 := 0
	divisor2 := 8
	divisor3 := 5

	printAll(SafeDivideAll(nums, float64(divisor1)))
	printAll(SafeDivideAll(nums, float64(divisor2)))
	printAll(SafeDivideAll(nums, float64(divisor3)))
}
