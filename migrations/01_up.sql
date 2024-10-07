CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,          -- Уникальный идентификатор пользователя
    login VARCHAR(255) NOT NULL,    -- Логин пользователя
    password VARCHAR(255) NOT NULL, -- Хэш пароля
    created_at TIMESTAMP NOT NULL DEFAULT NOW() -- Дата создания пользователя
);

CREATE TABLE IF NOT EXISTS places
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL, -- название места
    image VARCHAR(255) NOT NULL, -- путь к картинке
    description TEXT NOT NULL, -- описание места
    rating INT NOT NULL, -- рейтинг места
    numberOfReviews INT NOT NULL, -- количество отзывов
    address VARCHAR(255) NOT NULL, -- адрес места
    city VARCHAR(255) NOT NULL, -- город, где находится место
    phoneNumber VARCHAR(255), -- номер телефона
    category VARCHAR(255) NOT NULL -- категория места
);