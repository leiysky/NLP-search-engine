package main

import (
	"encoding/json"

	"github.com/syndtr/goleveldb/leveldb"
)

type Index struct {
	Filepath string `json:"filepath"`
	Count    int    `json:"count"`
}

type IndexBuilder struct {
	db *leveldb.DB
}

func NewIndexBuilder() (*IndexBuilder, error) {
	db, err := leveldb.OpenFile("data/index", nil)
	if err != nil {
		return nil, err
	}
	b := &IndexBuilder{
		db: db,
	}
	return b, nil
}

func (b *IndexBuilder) Get(token string) (*Index, error) {
	ret, err := b.db.Has([]byte(token), nil)
	if err != nil || !ret {
		return nil, err
	}
	res, err := b.db.Get([]byte(token), nil)
	if err != nil {
		return nil, err
	}
	index := &Index{}
	err = json.Unmarshal(res, index)
	if err != nil {
		return nil, err
	}
	return index, nil
}

func (b *IndexBuilder) Put(token string, index *Index) error {
	v, err := json.Marshal(index)
	if err != nil {
		return err
	}
	err = b.db.Put([]byte(token), v, nil)
	if err != nil {
		return err
	}
	return nil
}
