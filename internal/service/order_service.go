package main

/*
содержит " CreateOrder, GetOrder, ListOrders, CancelOrder "
вся логика по заказам:
- проверка статусов
- атомарное резервирование товаров
- расчет суммы

самый сложный файл - здесь будет транзакция с " SELECT ... FOR UPDATE "
*/
