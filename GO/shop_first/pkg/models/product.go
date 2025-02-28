package models

type Product struct {
	ID                    int
	Name                  string
	Price                 float64
	ImgURL                string
	AvailableQuantity     int
	Description           string
	Description_invisible string
}

/*
type Product struct {
	ID                    int     `json"ID"`
	Name                  string  `json"Name"`
	Price                 float64 `json"ID"`
	ImgURL                string  `json"ImgURL"`
	AvailableQuantity     int     `json"AvailableQuatity"`
	Description           string  `json"Description"`
	Description_invisible string  `json"Description_invisible"`
}
*/
