package order

import (
	"database/sql"
	"fmt"

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

func GetOrders(userID int64) (map[string]float64, error) {
	query := `SELECT product_name, quantity FROM "order" WHERE user_id = $1`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %v", err)
	}
	defer rows.Close()

	order := make(map[string]float64)
	for rows.Next() {
		var productName string
		var quantity float64
		if err := rows.Scan(&productName, &quantity); err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки: %v", err)
		}
		order[productName] = quantity
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при итерации по строкам: %v", err)
	}

	return order, nil
}
