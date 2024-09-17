package Math
// Calculate the Medin of the data
func Medin(data []float64) float64 {
	if len(data)%2 != 0 {
		return data[len(data)/2]
	} else {
		var c float64
		c = (data[len(data)/2] + data[(len(data)/2)-1]) / 2
		return c
	}

}
