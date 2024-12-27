// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Translation -.
type Translation struct {
	Source      string `json:"source"       example:"auto"`
	Destination string `json:"destination"  example:"en"`
	Original    string `json:"original"     example:"текст для перевода"`
	Translation string `json:"translation"  example:"text for translation"`
}
type Product struct {
	ID          string `json:"id" example:"1"`
	Name        string `json:"name" example:"Darius"`
	Description string `json:"description" example:"A great product"`
	Price       int    `json:"price" example:"100"`
	CreatedAt   string `json:"created_at" example:"2020-01-01"`
	UpdatedAt   string `json:"updated_at" example:"2020-01-01"`
}

type ProductHistory struct {
	ID          int    `json:"id" example:"1"`
	ProductID   int    `json:"product_id" example:"1"`
	Name        string `json:"name" example:"Darius"`
	Description string `json:"description" example:"A great product"`
	Price       int    `json:"price" example:"100"`
	ValidFrom   string `json:"valid_from" example:"2020-01-01"`
	ValidTo     string `json:"valid_to" example:"2020-01-01"`
	CreatedAt   string `json:"created_at" example:"2020-01-01"`
}

type ProductMaxValue struct {
	Price    int    `json:"price"`
	Duration string `json:"duration" example:"10"`
}

type TimeDiff struct {
	ValidFrom string `json:"valid_from"`
	ValidTo   string `json:"valid_to"`
	Price     int    `json:"price"`
}

type ReferenceDate struct {
	DateTime string `json:"date_time" example:"2020-01-01"`
}
