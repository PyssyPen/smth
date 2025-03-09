package main

import (
	"_pyssy_shop/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	// Обработчик для корневого URL
	http.HandleFunc("/", handlers.IndexHandler)

	// Обработчик для отдельной страницы продукта
	http.HandleFunc("/product/", handlers.ProductHandler)

	// Обработчик для формы добавления товара
	http.HandleFunc("/add_product", handlers.AddProductHandler)

	// Обработчик для обработки отправки формы добавления товара
	http.HandleFunc("/submit_product", handlers.SubmitProductHandler)

	// Обработчик для формы редактирования товара
	http.HandleFunc("/edit_product/", handlers.EditProductHandler)

	// Обработчик для обработки отправки формы редактирования товара
	http.HandleFunc("/update_product", handlers.UpdateProductHandler)

	// Обработчик для статических файлов (изображения и CSS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./pkg/static"))))

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
