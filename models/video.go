package models

type video struct{
	Id uint32 `gorm:"primary_key"`
	Url string
	Size uint32
	Uid uint32
	Title string
	CreateTime string
	Status uint8
}