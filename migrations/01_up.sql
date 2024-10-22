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

CREATE TABLE IF NOT EXISTS places
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL, -- название места
    image VARCHAR(255) NOT NULL, -- путь к картинке
    description TEXT NOT NULL -- описание места
);

CREATE TABLE IF NOT EXISTS trips_places ( --таблица для сопоставления поездки и достопримечательности, которая в нее входит
    id SERIAL PRIMARY KEY, 
    trip_id INT NOT NULL, -- Идентификатор поездки
    place_id INT NOT NULL, -- Идентификатор города
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания записи
    FOREIGN KEY (trip_id) REFERENCES trips(id) ON DELETE CASCADE,
    FOREIGN KEY (place_id) REFERENCES places(id) ON DELETE CASCADE
);
