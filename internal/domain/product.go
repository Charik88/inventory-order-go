package domain

/*
содержит ' type Product struct {ID, Name, Price, Stock...}'
одна сущность = один файл. легко найти и легко тестировать.
*/

import "time"

// Product - доменная модель товара (основная модель, содержит все
// поля, включая те, что генерирует база - id, created_at, updated_at)
type Product struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

/*
CreateProductRequest - DTO для создания товара
содержит только те поля, которые приходят от клиента
*/

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

/*
UpdateStockRequest - DTO для обновления остатка
*/

type UpdateStockRequest struct {
	Stock int `json:"stock"`
}
