package Math

// Calculate the variance of the data using the average
func Variance(data []float64, average float64) float64 {
	var conur float64
	for i := 0; i < len(data); i++ {
		conur += (data[i] - average) * (data[i] - average)
	}
	return conur / float64(len(data))
}
