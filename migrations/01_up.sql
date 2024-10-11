CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,          -- Уникальный идентификатор пользователя
    login VARCHAR(255) NOT NULL,    -- Логин пользователя
    password VARCHAR(255) NOT NULL, -- Хэш пароля
    created_at TIMESTAMP NOT NULL DEFAULT NOW() -- Дата создания пользователя
);

CREATE TABLE IF NOT EXISTS reviews
(
    id SERIAL PRIMARY KEY,                 -- Уникальный идентификатор отзыва
    user_id INT REFERENCES users(id),      -- Идентификатор пользователя, который оставил отзыв
    place_id INT REFERENCES places(id),    -- Идентификатор места, к которому относится отзыв
    rating INT NOT NULL CHECK (rating BETWEEN 1 AND 5), -- Рейтинг места (от 1 до 5)
    review_text TEXT,                      -- Текст отзыва
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания отзыва
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_place FOREIGN KEY (place_id) REFERENCES places(id) ON DELETE CASCADE
);