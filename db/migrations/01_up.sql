CREATE TABLE IF NOT EXISTS "user"
(
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,          -- Уникальный идентификатор пользователя
    login TEXT UNIQUE NOT NULL,    -- Логин пользователя
    email TEXT NOT NULL, -- Email пользователя
    password TEXT NOT NULL, -- Хэш пароля
    avatarPath VARCHAR(255) DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),-- Дата создания пользователя
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS city
(
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,          -- Уникальный идентификатор города
    name TEXT NOT NULL,    -- Название города
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания города
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS place
(
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL, -- название места
    imagePath TEXT NOT NULL DEFAULT '', -- путь к картинке
    description TEXT NOT NULL DEFAULT '', -- описание места
    rating INT NOT NULL DEFAULT 0, -- рейтинг места
    numberOfReviews INT NOT NULL DEFAULT 0, -- количество отзывов
    address TEXT NOT NULL DEFAULT '', -- адрес места
    cityId INT NOT NULL, -- город, где находится место
    phoneNumber TEXT DEFAULT '', -- номер телефона
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (cityId) REFERENCES city(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS review
(
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,                 -- Уникальный идентификатор отзыва
    user_id INT REFERENCES "user"(id) NOT NULL,      -- Идентификатор пользователя, который оставил отзыв
    place_id INT REFERENCES place(id) NOT NULL,    -- Идентификатор места, к которому относится отзыв
    rating INT NOT NULL CHECK (rating BETWEEN 1 AND 5) DEFAULT 0, -- Рейтинг места (от 1 до 5)
    review_text TEXT DEFAULT '',                      -- Текст отзыва
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания отзыва
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE,
    CONSTRAINT fk_place FOREIGN KEY (place_id) REFERENCES place(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS category
(
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS place_category
(
     place_id INT NOT NULL,
    category_id INT NOT NULL,
    PRIMARY KEY(place_id, category_id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (place_id) REFERENCES place(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS trip
(
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,          -- Уникальный идентификатор поездки
    name TEXT NOT NULL, -- Название поездки
    description TEXT DEFAULT '', -- Описание поездки
    city_id INTEGER NOT NULL, -- Направление поездки
    start_date DATE NOT NULL, -- Дата начала поездки
    end_date DATE NOT NULL, -- Дата окончания поездки
    private BOOLEAN NOT NULL DEFAULT TRUE, -- Кому видна поездка (всем или выбранным пользователям)
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания поездки
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id INTEGER NOT NULL, -- Идентификатор пользователя-создателя поездки
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE,
    CONSTRAINT fk_city FOREIGN KEY (city_id) REFERENCES city(id) ON DELETE CASCADE
 );

CREATE TABLE IF NOT EXISTS trip_place ( --таблица для сопоставления поездки и достопримечательности, которая в нее входит
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY, 
    trip_id INT NOT NULL, -- Идентификатор поездки
    place_id INT NOT NULL, -- Идентификатор города
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания записи
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (trip_id) REFERENCES trip(id) ON DELETE CASCADE,
    FOREIGN KEY (place_id) REFERENCES place(id) ON DELETE CASCADE
);

COPY city(name)
    FROM '/docker-entrypoint-initdb.d/cities.csv'
    WITH (FORMAT csv, HEADER true);

COPY placs(name, imagePath, description, rating, numberOfReviews, address, cityId, phoneNumber)
    FROM '/docker-entrypoint-initdb.d/places.csv'
    WITH (FORMAT csv, HEADER true,  DELIMITER ';');

COPY category(name)
    FROM '/docker-entrypoint-initdb.d/categories.csv'
    WITH (FORMAT csv, HEADER true);

COPY place_category(place_id,category_id)
    FROM '/docker-entrypoint-initdb.d/places_categories.csv'
    WITH (FORMAT csv, HEADER true, DELIMITER ',');
