package model

type Goods struct {
	Id     int `gorm:"primaryKey"`
	Name   string
	Url    string
	UrlIMG string
	Price  int
}
