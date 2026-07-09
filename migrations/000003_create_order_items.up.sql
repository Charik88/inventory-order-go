/*
создает таблицу orders_items (состав заказа)
*/

CREATE TABLE IF NOT EXISTS order_items(
    order_id BEGINT NOT NULL REFERENCES orders(id) ON DELETE CASCADE, // совпадает с bigserial из orders.id. внешний ключ гарантирует целостность  данных, нельзя вставить заказ которого нет в таблице заказов
    product_id BEGINT NOT NULL REFERENCES products(id),// нельзя удалить товар вот есть в заказе
    quantity INT NOT NULL CHECK (quantity > 0), // кол-во ед товара в заказе
    price DECIMAL (12,2) NOT NULL, // цена товара на момент заказа
    PRIMARY KEY (order_id, product_id) // (одна и таже пара) нельзя добавить один и тот же товар в од ин заказ дважды
);