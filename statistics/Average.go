package Math

// Calculate the average of the data
func Average(datasep []float64) float64 {
	var countr float64
	for i := 0; i < len(datasep); i++ {
		countr += datasep[i]

	}

	return (countr / float64(len(datasep)))

}
