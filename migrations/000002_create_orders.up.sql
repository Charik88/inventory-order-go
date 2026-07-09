/*
создает таблицу orders (заказы)
*/
CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY,
    status TEXT NOT NULL DEFAULT 'created'
                                  CHECK (status in ('created', 'cancelled')),
    created_at TIMESTAMTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status);
// индек отдельная структура данных кот ускоряет поиск по полю status
// можно и бех него, но если маленькие обьемы данных