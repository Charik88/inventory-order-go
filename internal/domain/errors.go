package domain

/*
содержит 'var ErrNotFound, ErrInsufficientStock, ErrConflict' - доменные ошибки
Все ошибки в одном месте. Сервисы возвращают именно их, хендлеры мапят в HTTP - статусы.
не придется искать "errors.New" по всему проекту.
*/

import "errors"

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("conflict")
	ErrValidation        = errors.New("validation error")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

/*
смыссловые ошибки, а не технические
Repository возвращает ErrNotFound, а hendler превращает ее в HTTP 404


ErrNotFound - товар или заказ не найлен (404)
ErrConflict - конфликт данных (дубликат) (409)
ErrValidation - неверные входные данные (400)
ErrInsufficientFunds - нехватает товара на складе (422)
*/
