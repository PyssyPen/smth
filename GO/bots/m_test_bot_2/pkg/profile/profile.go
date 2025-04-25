package profile

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

func GetProfile(userID int64) (map[string]string, error) {
	query := `SELECT username, phone_number, delivery_address FROM profiles WHERE user_id = $1`
	row := db.QueryRow(query, userID)

	var username, phoneNumber, deliveryAddress string
	if err := row.Scan(&username, &phoneNumber, &deliveryAddress); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("профиль не найден")
		}
		return nil, fmt.Errorf("ошибка при получении профиля: %v", err)
	}

	return map[string]string{
		"username":         username,
		"phone_number":     phoneNumber,
		"delivery_address": deliveryAddress,
	}, nil
}

func UpdateProfile(userID int64, username, phoneNumber, deliveryAddress string) error {
	query := `INSERT INTO profiles (user_id, username, phone_number, delivery_address)
              VALUES ($1, $2, $3, $4)
              ON CONFLICT (user_id) DO UPDATE
              SET username = COALESCE($2, username),
                  phone_number = COALESCE($3, phone_number),
                  delivery_address = COALESCE($4, delivery_address)`
	_, err := db.Exec(query, userID, username, phoneNumber, deliveryAddress)
	return err
}
