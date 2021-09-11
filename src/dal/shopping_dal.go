package dal

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"gogin/src/models"
)

type Dal struct {
}

func (Dal) GetAllItems(db *sqlx.DB) []models.Items {

	getAllQuery := "SELECT name, price, quantity from Items"

	items := []models.Items{}

	db.Select(&items, getAllQuery)

	return items
}

func (Dal) GetItemsByCategory(db *sqlx.DB, category string) ([]models.Items, error) {
	getAllQuery := `SELECT name, price, quantity from Items where category = "` + category + `"`

	items := []models.Items{}

	err := db.Select(&items, getAllQuery)

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (Dal) PostItems(db *sqlx.DB, item models.NewItems) error {
	getAllQuery := `INSERT INTO items (name, category, price, quantity) VALUE ("` + item.Name + `", "` + item.Category + `", ` + item.Price + `, ` + item.Quantity + `)`
	fmt.Println(getAllQuery)
	_, err := db.Exec(getAllQuery)

	if err != nil {
		return err
	}

	return nil
}
