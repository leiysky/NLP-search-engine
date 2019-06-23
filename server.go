package main

import (
	"io/ioutil"
	"net/url"
	"regexp"
	"sort"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	queries, err := url.QueryUnescape(c.Query("q"))
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	regex := regexp.MustCompile(`[\w]+`)
	query := regex.FindAllString(queries, -1)
	index := make(map[string][]Index)

	for _, q := range query {
		if _, ok := tokens[q]; ok {
			index[q] = tokens[q]
		}
	}

	filter := make(map[string]int)
	for _, v := range index {
		for _, idx := range v {
			filter[idx.Filepath] = filter[idx.Filepath] + 1
		}
	}

	var result Ranker
	for k, v := range filter {
		if v == len(query) {
			var score float64
			for _, t := range query {
				score += scores[t][k]
			}
			result = append(result, struct {
				filepath string
				score    float64
			}{k, score})
		}
	}

	sort.Sort(result)

	var ret []string
	for _, s := range result {
		buff, err := ioutil.ReadFile(s.filepath)
		if err != nil {
			panic(err)
		}
		ret = append(ret, string(buff))
	}

	c.JSON(200, ret)
}

type Ranker []struct {
	filepath string
	score    float64
}

func (r Ranker) Less(a, b int) bool {
	if r[a].score > r[b].score {
		return true
	}
	return false
}

func (r Ranker) Swap(a, b int) {
	r[a], r[b] = r[b], r[a]
}

func (r Ranker) Len() int {
	return len(r)
}
