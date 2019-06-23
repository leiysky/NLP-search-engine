package main

import "math"

const DocCount float64 = 193

// token -> filepath -> tfidf
func Tfidf(tokens map[string][]Index) map[string]map[string]float64 {
	ret := make(map[string]map[string]float64)
	for t, index := range tokens {
		ret[t] = make(map[string]float64)
		for _, idx := range index {
			f := idx.Filepath
			tf := float64(1) + math.Log10(float64(idx.Count))
			df := float64(len(index))
			idf := math.Log10(DocCount / df)
			w := (1 + math.Log10(tf)) * idf
			ret[t][f] = w
		}
	}
	return ret
}
