package main

import (
	"io/ioutil"
	"os"
	"regexp"

	"github.com/yanyiwu/gojieba"
)

type TokenStream struct {
	x *gojieba.Jieba
	s []string
}

func NewTokenStream() *TokenStream {
	ts := &TokenStream{
		x: gojieba.NewJieba(),
	}
	return ts
}

func (ts *TokenStream) Reset(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	buff, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	ts.s = ts.x.CutForSearch(string(buff), true)
	return nil
}

type TokenIterator struct {
	s []string
}

func NewTokenIterator(ts *TokenStream) *TokenIterator {
	return &TokenIterator{
		s: ts.s,
	}
}

func (it *TokenIterator) ForEach(cb func(token string)) {
	for _, s := range it.s {
		cb(s)
	}
}

func (it *TokenIterator) Get() []string {
	return it.s
}

func (it *TokenIterator) Map(cb func(token string) string) *TokenIterator {
	var res []string
	for _, s := range it.s {
		res = append(res, cb(s))
	}
	return &TokenIterator{
		s: res,
	}
}

func (it *TokenIterator) Filter(cb func(token string) bool) *TokenIterator {
	var res []string
	for _, s := range it.s {
		if cb(s) {
			res = append(res, s)
		}
	}
	return &TokenIterator{
		s: res,
	}
}

func isValidToken(token string) bool {
	regex := regexp.MustCompile(`^[\p{Han}]+|[\w]+$`)
	return regex.MatchString(token)
}
