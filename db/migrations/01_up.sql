CREATE TABLE IF NOT EXISTS "user"
(
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,          -- Уникальный идентификатор пользователя
    login TEXT UNIQUE NOT NULL,    -- Логин пользователя
    email TEXT UNIQUE NOT NULL, -- Email пользователя
    password_hash TEXT NOT NULL, -- Хэш пароля
    avatar_path TEXT DEFAULT '',
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
CREATE TABLE IF NOT EXISTS category
(
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT UNIQUE NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS place
(
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL, -- название места
    image_path TEXT NOT NULL DEFAULT '', -- путь к картинке
    description TEXT NOT NULL DEFAULT '', -- описание места
    rating DECIMAL(2,1) NOT NULL DEFAULT 0.0, -- рейтинг места
    address TEXT NOT NULL DEFAULT '', -- адрес места
    city_id INT NOT NULL, -- город, где находится место
    phone_number TEXT DEFAULT '', -- номер телефона
    number_of_reviews INTEGER NOT NULL DEFAULT 0,
    latitude DECIMAL(7,4),
    longitude DECIMAL(7,4), 
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (city_id) REFERENCES city(id) ON DELETE CASCADE,
    CONSTRAINT check_place_rating CHECK (rating BETWEEN 0.0 AND 5.0), -- Ограничение для rating
    CONSTRAINT uq_place_name_city UNIQUE (name, city_id)
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
    CONSTRAINT fk_place FOREIGN KEY (place_id) REFERENCES place(id) ON DELETE CASCADE,
    CONSTRAINT uq_user_place_review UNIQUE (user_id, place_id)
);



CREATE TABLE IF NOT EXISTS place_category
(
    place_id INT NOT NULL,
    category_id INT NOT NULL,
    PRIMARY KEY(place_id, category_id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (place_id) REFERENCES place(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE,
    CONSTRAINT uq_place_category UNIQUE (place_id, category_id)
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
    FOREIGN KEY (place_id) REFERENCES place(id) ON DELETE CASCADE,
    CONSTRAINT uq_trip_place UNIQUE (trip_id, place_id)
);

CREATE TABLE trip_photo (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY, 
    trip_id INT REFERENCES trip(id) ON DELETE CASCADE,
    photo_path TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS survey (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    survey_text TEXT DEFAULT '',
    max_rating INT NOT NULL CHECK (max_rating > 0) DEFAULT 1
);

CREATE TABLE IF NOT EXISTS user_survey (
    survey_id INT NOT NULL,
    user_id INT NOT NULL,
    PRIMARY KEY(survey_id, user_id),
    rating INT NOT NULL CHECK (rating > 0) DEFAULT 1,
    FOREIGN KEY (survey_id) REFERENCES survey(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS sharing_token (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    trip_id INT NOT NULL, -- ID страницы
    token TEXT NOT NULL UNIQUE, -- Уникальный токен для ссылки
    sharing_option TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания записи
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP, -- Срок действия ссылки (если нужен)
    CONSTRAINT fk_page FOREIGN KEY (trip_id) REFERENCES trip(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_shared_trip (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    trip_id INT NOT NULL, -- ID страницы
    user_id INT NOT NULL, -- ID страницы
    sharing_option TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания записи
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_trip FOREIGN KEY (trip_id) REFERENCES trip(id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS achievement (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,             -- Название достижения
    icon_path TEXT NOT NULL,        -- Ссылка на изображение
    created_at TIMESTAMP NOT NULL DEFAULT NOW()      -- Дата создания записи
);

CREATE TABLE IF NOT EXISTS user_achievement (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL,                   -- ID пользователя
    achievement_id INT NOT NULL,            -- ID достижения
    FOREIGN KEY (achievement_id) REFERENCES achievement (id) ON DELETE CASCADE
);

COPY city(name)
    FROM '/docker-entrypoint-initdb.d/cities.csv'
    WITH (FORMAT csv, HEADER true);

COPY place(name, image_path, description, rating, address, city_id, phone_number, latitude, longitude)
    FROM '/docker-entrypoint-initdb.d/places.csv'
    WITH (FORMAT csv, HEADER true,  DELIMITER ';');

COPY category(name)
    FROM '/docker-entrypoint-initdb.d/categories.csv'
    WITH (FORMAT csv, HEADER true);

COPY place_category(place_id,category_id)
    FROM '/docker-entrypoint-initdb.d/places_categories.csv'
    WITH (FORMAT csv, HEADER true, DELIMITER ',');

COPY achievement(name, icon_path)
    FROM '/docker-entrypoint-initdb.d/achievements.csv'
    WITH (FORMAT csv, HEADER true, DELIMITER ';');
