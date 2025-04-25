package cart

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

func CreateOrder(userID int64) error {
	fmt.Printf("Создание заказа для UserID: %d\n", userID)

	// Получение всех записей из корзины для данного пользователя
	query := `SELECT product_name, quantity FROM cart WHERE user_id = $1`
	rows, err := db.Query(query, userID)
	if err != nil {
		fmt.Printf("Ошибка при получении данных из корзины: %v\n", err)
		return err
	}
	defer rows.Close()

	// Вставка записей в таблицу "order"
	for rows.Next() {
		var productName string
		var quantity float64
		if err := rows.Scan(&productName, &quantity); err != nil {
			fmt.Printf("Ошибка при сканировании строки: %v\n", err)
			return err
		}
		insertQuery := `INSERT INTO "order" (user_id, product_name, quantity) VALUES ($1, $2, $3)`
		if _, err := db.Exec(insertQuery, userID, productName, quantity); err != nil {
			fmt.Printf("Ошибка при вставке данных в таблицу order: %v\n", err)
			return err
		}
	}

	// Очистка корзины
	if err := ClearCart(userID); err != nil {
		fmt.Printf("Ошибка при очистке корзины: %v\n", err)
		return err
	}

	return nil
}

// Функция для очистки корзины пользователя
func ClearCart(userID int64) error {
	fmt.Printf("Очистка корзины для UserID: %d\n", userID)
	query := `DELETE FROM cart WHERE user_id = $1`
	_, err := db.Exec(query, userID)
	if err != nil {
		fmt.Printf("Ошибка при очистке корзины: %v\n", err)
	}
	return err
}
