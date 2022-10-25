package model

type Menu struct {
	Id string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Price int `json:"price" binding:"required"`
}

func NewMenu(id, name string, price int) Menu {
	return Menu{
		Id:      id,
		Name:    name,
		Price: price,
	}

}