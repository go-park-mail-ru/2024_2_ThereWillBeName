CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,          -- Уникальный идентификатор пользователя
    login VARCHAR(255) NOT NULL,    -- Логин пользователя
    password VARCHAR(255) NOT NULL, -- Хэш пароля
    created_at TIMESTAMP NOT NULL DEFAULT NOW() -- Дата создания пользователя
);

CREATE TABLE IF NOT EXISTS cities
(
    id SERIAL PRIMARY KEY,          -- Уникальный идентификатор города
    name VARCHAR(255) NOT NULL,    -- Название города
    created_at TIMESTAMP NOT NULL DEFAULT NOW() -- Дата создания города
);

CREATE TABLE IF NOT EXISTS places
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL, -- название места
    imagePath VARCHAR(255) NOT NULL, -- путь к картинке
    description TEXT NOT NULL, -- описание места
    rating INT NOT NULL, -- рейтинг места
    numberOfReviews INT NOT NULL, -- количество отзывов
    address VARCHAR(255) NOT NULL, -- адрес места
    cityId INT NOT NULL, -- город, где находится место
    phoneNumber VARCHAR(10), -- номер телефона
    FOREIGN KEY (cityId) REFERENCES cities(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS categories
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS places_categories
(
    place_id INT NOT NULL,
    category_id INT NOT NULL,
    PRIMARY KEY(place_id, category_id),
    FOREIGN KEY (place_id) REFERENCES places(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);
