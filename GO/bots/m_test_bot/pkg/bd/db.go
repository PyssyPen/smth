package db

import (
	"database/sql"
	"m_test_bot/pkg/models"

	_ "github.com/lib/pq"
)

// ConnectDB функция для подключения к базе данных
func ConnectDB() (*sql.DB, error) {
	connStr := "user=postgres dbname=pyssyshop sslmode=disable password=qwerty"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetProducts функция для получения данных о товарах из базы данных pyssyshop, таблицы products
func GetProducts(db *sql.DB) ([]models.Product, error) {
	rows, err := db.Query("SELECT id, name, price, img_url, available_quantity, description, description_invisible FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.ImgURL, &p.AvailableQuantity, &p.Description, &p.Description_invisible); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// GetProductByID функция для получения данных о товаре по ID
func GetProductByID(db *sql.DB, id int) (models.Product, error) {
	var p models.Product
	err := db.QueryRow("SELECT id, name, price, img_url, available_quantity, description, description_invisible FROM products WHERE id=$1", id).Scan(&p.ID, &p.Name, &p.Price, &p.ImgURL, &p.AvailableQuantity, &p.Description, &p.Description_invisible)
	if err != nil {
		return p, err
	}
	return p, nil
}

// AddProduct функция для добавления нового товара в базу данных
func AddProduct(db *sql.DB, product models.Product) error {
	_, err := db.Exec("INSERT INTO products (name, price, img_url, available_quantity, description, description_invisible) VALUES ($1, $2, $3, $4, $5, $6)",
		product.Name, product.Price, product.ImgURL, product.AvailableQuantity, product.Description, product.Description_invisible)
	return err
}

// UpdateProduct функция для обновления товара в базе данных
func UpdateProduct(db *sql.DB, product models.Product) error {
	_, err := db.Exec("UPDATE products SET name=$1, price=$2, img_url=$3, available_quantity=$4, description=$5, description_invisible=$6 WHERE id=$7",
		product.Name, product.Price, product.ImgURL, product.AvailableQuantity, product.Description, product.Description_invisible, product.ID)
	return err
}
