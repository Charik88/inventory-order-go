package domain

/*
содержит  " type Pagination struct { Page, Limit} " и " type PaginatedResult struct {Data, Total}"
сквозная утилита. используется и продуктами и заказами. выносим отдельно чтобы не дублировать.
*/
// Pagination - параметры пагинации от клиетна
// разбивка большого списка на страницы.
//Вместо того чтобы отдавать клиенту все 10 000 товаров одним куском, сервер отдаёт по 20 штук за раз.

type Pagination struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

// Offset вычисляет смещение для SQL-запроса

func (p Pagination) Offset() int {
	if p.Page < 1 {
		p.Page = 1
	}
	return (p.Page - 1) * p.Size
}

// PaginatedResult - обертка для ответа со списком

type PaginationResult[T any] struct {
	Items []T `json:"items"`
	Total int `json:"total"`
	Page  int `json:"page"`
	Size  int `json:"size"`
}
