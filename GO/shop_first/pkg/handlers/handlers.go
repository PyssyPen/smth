package handlers

import (
	"html/template"
	"net/http"
	"pyssyshop/pkg/db"
	"pyssyshop/pkg/models"
	"strconv"
)

// IndexHandler обработчик для корневого URL
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	dbConn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	products, err := db.GetProducts(dbConn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("pkg/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, products)
}

// ProductHandler обработчик для отдельной страницы продукта
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/product/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	dbConn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	product, err := db.GetProductByID(dbConn, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("pkg/templates/product.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, product)
}

// AddProductHandler обработчик для формы добавления товара
func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pkg/templates/add_product.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

// SubmitProductHandler обработчик для обработки отправки формы добавления товара
func SubmitProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			http.Error(w, "Invalid price", http.StatusBadRequest)
			return
		}
		imgURL := r.FormValue("img_url")
		availableQuantity, err := strconv.Atoi(r.FormValue("available_quantity"))
		if err != nil {
			http.Error(w, "Invalid available quantity", http.StatusBadRequest)
			return
		}
		description := r.FormValue("description")
		description_invisible := r.FormValue("description_invisible")

		newProduct := models.Product{
			Name:                  name,
			Price:                 price,
			ImgURL:                imgURL,
			AvailableQuantity:     availableQuantity,
			Description:           description,
			Description_invisible: description_invisible,
		}

		dbConn, err := db.ConnectDB()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dbConn.Close()

		err = db.AddProduct(dbConn, newProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// EditProductHandler обработчик для формы редактирования товара
func EditProductHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/edit_product/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	dbConn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	product, err := db.GetProductByID(dbConn, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("pkg/templates/edit_product.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, product)
}

// UpdateProductHandler обработчик для обработки отправки формы редактирования товара
func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			http.Error(w, "Invalid price", http.StatusBadRequest)
			return
		}
		imgURL := r.FormValue("img_url")
		availableQuantity, err := strconv.Atoi(r.FormValue("available_quantity"))
		if err != nil {
			http.Error(w, "Invalid available quantity", http.StatusBadRequest)
			return
		}
		description := r.FormValue("description")
		description_invisible := r.FormValue("description_invisible")

		updatedProduct := models.Product{
			ID:                    id,
			Name:                  name,
			Price:                 price,
			ImgURL:                imgURL,
			AvailableQuantity:     availableQuantity,
			Description:           description,
			Description_invisible: description_invisible,
		}

		dbConn, err := db.ConnectDB()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dbConn.Close()

		err = db.UpdateProduct(dbConn, updatedProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/product/"+strconv.Itoa(id), http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
