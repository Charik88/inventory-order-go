package main

/*
содержит 'var ErrNotFound, ErrInsufficientStock, ErrConflict' - доменные ошибки
Все ошибки в одном месте. Сервисы возвращают именно их, хендлеры мапят в HTTP - статусы.
не придется искать "errors.New" по всему проекту.
*/
