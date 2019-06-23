package main

import "github.com/gin-gonic/gin"

var tokens map[string][]Index
var scores map[string]map[string]float64

func init() {
	p, err := NewParser()
	if err != nil {
		panic(err)
	}
	tokens, err = p.Parse()
	if err != nil {
		panic(err)
	}
	scores = Tfidf(tokens)
}

func main() {
	r := gin.Default()
	r.GET("/search", Search)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
