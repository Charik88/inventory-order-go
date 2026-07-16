package domain

// OrderStatus - статус заказа (enum через константы)

import "time"

type OrderStatus string

const (
	OrderStatusCreated   OrderStatus = "created"
	OrderStatusCancelled OrderStatus = "cancelled"
)

// OrderItem -  элемент заказа (товар + кол-во)

type OrderItem struct {
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

// Order - доменная модель заказа

type Order struct {
	ID        int64       `json:"id"`
	Status    OrderStatus `json:"status"`
	Items     []OrderItem `json:"items,omitempty"` // omitempty - если заказы списком, без состава, поле не попадет в JSON
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// CreateOrderRequest - DTO ля создания заказа.
// DTO: клиент присылает массив items - какие товары и в каком кол-ве
type CreateOrderRequest struct {
	Items []OrderItem `json:"items"`
}
