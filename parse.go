package main

import (
	"os"
	"path/filepath"
)

type Parser struct {
	b  *IndexBuilder
	ts *TokenStream
}

func NewParser() (*Parser, error) {
	// b, err := NewIndexBuilder()
	// if err != nil {
	// 	return nil, err
	// }
	return &Parser{
		// b:  b,
		ts: NewTokenStream(),
	}, nil
}

func (p *Parser) Parse() (map[string][]Index, error) {
	var files []string
	tokens := make(map[string][]Index)
	err := filepath.Walk("page", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		p.ts.Reset(f)
		it := NewTokenIterator(p.ts)
		it.Filter(isValidToken).ForEach(func(token string) {
			index := tokens[token]
			for i, idx := range index {
				if idx.Filepath == f {
					index[i].Count++
					tokens[token] = index
					return
				}
			}
			index = append(index, Index{f, 1})
			tokens[token] = index
		})
	}
	// for _, idx := range tokens {
	// 	fmt.Println(idx)
	// }
	return tokens, nil
}
