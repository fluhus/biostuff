/* ****************************************************************************
 * Statistic functions.
 * ****************************************************************************/

// Statistic functions.
package stat

import "math"

// Returns the sum of the values in a sample.
func Sum(sample []float64) float64 {
	s := float64(0)
	for _,v := range sample {s += v}
	return s
}

// Returns the mean value of the sample.
func Mean(sample []float64) float64 {
	return Sum(sample) / float64(len(sample))
}

// Returns the covariance of the 2 samples.
// Returns NaN for mismatching sample sizes, or 0 size samples.
func Covariance(sample1, sample2 []float64) float64 {
	if len(sample1) != len(sample2) ||
		len(sample1) == 0 ||
		len(sample2) == 0 {return math.NaN()}

	m1 := Mean(sample1)
	m2 := Mean(sample2)
	cov := float64(0)
	for i := range sample1 {
		cov += (sample1[i] - m1) * (sample2[i] - m2)
	}
	cov /= float64(len(sample1))

	return cov
}

// Returns the variance of the sample.
func Variance(sample []float64) float64 {
	return Covariance(sample, sample)
}

// Returns the standard deviation of the sample.
func Std(sample []float64) float64 {
	return math.Sqrt(Variance(sample))
}

// Returns the correlation between the samples.
func Correlation(sample1, sample2 []float64) float64 {
	return Covariance(sample1, sample2) / Std(sample1) / Std(sample2)
}

// Returns the minimal element in the sample.
func Min(sample []float64) float64 {
	if len(sample) == 0 {return math.NaN()}

	min := sample[0]
	for _,s := range sample {
		if s < min {min = s}
	}

	return min
}

// Returns the maximal element in the sample.
func Max(sample []float64) float64 {
	if len(sample) == 0 {return math.NaN()}

	max := sample[0]
	for _,s := range sample {
		if s > max {max = s}
	}

	return max
}

// Returns the span of the sample (max - min).
func Span(sample []float64) float64 {
	return Max(sample) - Min(sample)
}

// Returns the entropy for the given distribution.
// The distribution does not have to sum up to 1, for it will be normalized
// anyway.
func Entropy(distribution []float64) float64 {
	// Sum of the distribution
	sum := Sum(distribution)

	// Go over each bucket
	result := 0.0
	for _,v := range distribution {
		// Negative values are not allowed
		if v < 0.0 {
			return math.NaN()
		}

		// Ignore zeros
		if v == 0.0 {
			continue
		}

		// Probability
		p := v / sum

		// Entropy
		result -= p * math.Log2(p)
	}

	return result
}



