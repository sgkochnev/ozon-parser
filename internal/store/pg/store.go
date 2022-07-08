package pg

import (
	"ozon-parser/model"
)

type Inserter interface {
	Insert(goods ...model.Goods) []int
}

type store struct {
	db *DB
}

func New() (*store, error) {
	db, err := Dial()
	if err != nil {
		return nil, err
	}
	return &store{db: db}, nil
}

func (s *store) Insert(goods ...model.Goods) ([]int, error) {
	if err := s.db.Create(&goods).Error; err != nil {
		return nil, err
	}
	id := make([]int, len(goods))
	for i, v := range goods {
		id[i] = v.Id
	}
	return id, nil
}
