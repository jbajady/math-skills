package Math

import "strconv"
// Changeintflot converts a slice of strings into a slice of floats
func Changeintflot(data []string) ([]float64, error) {
	sl := []float64{}
	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			continue
		} else {
			numbr, err := strconv.Atoi(data[i])
			if err != nil {
				return nil, err
			}
			sl = append(sl, float64(numbr))
		}
	}
	return sl, nil
}
