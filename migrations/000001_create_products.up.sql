/*
создает таблицу products (товары)
*/
CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY, // код товара
    name TEXT NOT NULL, // название товара (не должно бфть пустым)
    price DECIMAL (12,2) NOT NULL CHECK (price >= 0), // цена с копейками, проверяем чтобы цена не была меньше 0
    stock INT NOT NULL DEFAULT 0 CHECK (stock >= 0), // кол-во товара не должно быть отрицательным
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (), // дата и время
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()  // дата последнего обновления
);