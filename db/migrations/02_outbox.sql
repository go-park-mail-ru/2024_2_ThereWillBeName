CREATE TABLE IF NOT EXISTS outbox (
    id SERIAL PRIMARY KEY,
    event_type TEXT NOT NULL,       -- Тип события
    payload TEXT NOT NULL,         -- Данные события
    "status" TEXT DEFAULT 'pending',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    processed_at TIMESTAMP          -- Время обработки события
);
