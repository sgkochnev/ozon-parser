package app

import (
	"fmt"
	"ozon-parser/internal/store/pg"
)

const baseURL = "https://www.ozon.ru"

const fullURL = baseURL + "/category/moloko-9283/"

// html для теста
// const staticHTML = "ozon.html"

func Run() error {
	store, err := pg.New()
	if err != nil {
		return err
	}

	goods, err := parseOzon(fullURL)
	if err != nil {
		return err
	}
	if len(goods) == 0 {
		fmt.Println("0 goods")
		return nil
	}
	_, err = store.Insert(goods...)
	if err != nil {
		return err
	}
	return nil
}
