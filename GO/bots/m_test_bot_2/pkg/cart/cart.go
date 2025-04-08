package cart

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(connStr string) error {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	return db.Ping()
}

func AddToCart(userID int64, productName string, quantity float64) error {
	query := `INSERT INTO cart (user_id, product_name, quantity) VALUES ($1, $2, $3)
              ON CONFLICT (user_id, product_name) DO UPDATE SET quantity = cart.quantity + $3`
	_, err := db.Exec(query, userID, productName, quantity)
	return err
}

func RemoveFromCart(userID int64, productName string) error {
	query := `DELETE FROM cart WHERE user_id = $1 AND product_name = $2`
	_, err := db.Exec(query, userID, productName)
	return err
}

func GetCart(userID int64) (map[string]float64, error) {
	query := `SELECT product_name, quantity FROM cart WHERE user_id = $1`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cart := make(map[string]float64)
	for rows.Next() {
		var productName string
		var quantity float64
		if err := rows.Scan(&productName, &quantity); err != nil {
			return nil, err
		}
		cart[productName] = quantity
	}
	return cart, nil
}
