CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,          -- Уникальный идентификатор пользователя
    login VARCHAR(255) NOT NULL UNIQUE,    -- Логин пользователя
    avatarPath VARCHAR(255) DEFAULT '',
    email TEXT NOT NULL,
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
    phoneNumber TEXT, -- номер телефона
    FOREIGN KEY (cityId) REFERENCES cities(id) ON DELETE CASCADE
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

CREATE TABLE IF NOT EXISTS trips
(
    id SERIAL PRIMARY KEY,          -- Уникальный идентификатор поездки
    name VARCHAR(255) NOT NULL, -- Название поездки
    description VARCHAR(1000), -- Описание поездки
    city_id INTEGER NOT NULL, -- Направление поездки
    start_date DATE, -- Дата начала поездки
    end_date DATE, -- Дата окончания поездки
    private BOOLEAN DEFAULT TRUE, -- Кому видна поездка (всем или выбранным пользователям)
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания поездки
    user_id INTEGER NOT NULL, -- Идентификатор пользователя-создателя поездки
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_city FOREIGN KEY (city_id) REFERENCES cities(id) ON DELETE CASCADE
 );

CREATE TABLE IF NOT EXISTS trips_places ( --таблица для сопоставления поездки и достопримечательности, которая в нее входит
    id SERIAL PRIMARY KEY, 
    trip_id INT NOT NULL, -- Идентификатор поездки
    place_id INT NOT NULL, -- Идентификатор города
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания записи
    FOREIGN KEY (trip_id) REFERENCES trips(id) ON DELETE CASCADE,
    FOREIGN KEY (place_id) REFERENCES places(id) ON DELETE CASCADE
);

COPY cities(name)
    FROM '/docker-entrypoint-initdb.d/cities.csv'
    WITH (FORMAT csv, HEADER true);

COPY places(name, imagePath, description, rating, numberOfReviews, address, cityId, phoneNumber)
    FROM '/docker-entrypoint-initdb.d/places.csv'
    WITH (FORMAT csv, HEADER true,  DELIMITER ';');

COPY categories(name)
    FROM '/docker-entrypoint-initdb.d/categories.csv'
    WITH (FORMAT csv, HEADER true);

COPY places_categories(place_id,category_id)
    FROM '/docker-entrypoint-initdb.d/places_categories.csv'
    WITH (FORMAT csv, HEADER true, DELIMITER ',');
