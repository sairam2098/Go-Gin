package services

import (
	"fmt"

	"gogin/src/dal"
	"gogin/src/database"
	"gogin/src/models"
)

type ShoppingService struct {
}

var shoppingDal dal.Dal = dal.Dal{}
var da database.DataAccessor = database.DataAccessor{}

func (ShoppingService) GetItems() []models.Items {
	db, err := da.Connect()
	defer da.Close(db)
	if err != nil {
		msg := err.Error()
		fmt.Print(msg)
		return nil
	}

	items := shoppingDal.GetAllItems(db)
	return items
}

func (ShoppingService) GetItemsByCategory(category string) ([]models.Items, error) {
	db, err := da.Connect()
	defer da.Close(db)
	if err != nil {
		msg := err.Error()
		fmt.Print(msg)
		return nil, err
	}

	items, err := shoppingDal.GetItemsByCategory(db, category)
	return items, err
}
