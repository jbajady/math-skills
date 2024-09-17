package main

import (
	"math"
	"os"
	"sort"
	"strings"

	"fmt"
	Math "math-skills/statistics"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		return
	}
	data, er := os.ReadFile(os.Args[1])
	if er != nil {
		fmt.Println("Error reading file")
		return
	}
	if len(data) == 0 {
		fmt.Println("Error file is empty")
		return
	}
	datasep := strings.Split(string(data), "\n")
	datase, err := Math.Changeintflot(datasep)
	if err != nil || len(datase) == 0 {
		fmt.Println("Error converting data to float")
		return
	}
	sort.Float64s(datase) // Sort the slice of floats in ascending order
	average := Math.Average(datase)
	medin := Math.Medin(datase)
	variance := Math.Variance(datase, Math.Average(datase))
	stabdarddevviation := Math.StandardDeviation(variance)

	fmt.Println("Average: ", math.Round(average))
	fmt.Println("Medin: ", math.Round(medin))
	fmt.Println("Variance: ", math.Round(variance))
	fmt.Println("StabdardDevviation: ", math.Round(stabdarddevviation))

}
