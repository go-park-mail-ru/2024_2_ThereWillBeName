CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,          -- Уникальный идентификатор пользователя
    login VARCHAR(255) NOT NULL,    -- Логин пользователя
    password VARCHAR(255) NOT NULL, -- Хэш пароля
    created_at TIMESTAMP NOT NULL DEFAULT NOW() -- Дата создания пользователя
);

CREATE TABLE IF NOT EXISTS trips
(
    id SERIAL PRIMARY KEY,          -- Уникальный идентификатор поездки
    name VARCHAR(255) NOT NULL, -- Название поездки
    description VARCHAR(1000), -- Описание поездки
    destination VARCHAR(255), -- Направление поездки
    start_date VARCHAR(10), -- Дата начала поездки
    end_date VARCHAR(10), -- Дата окончания поездки
    private BOOLEAN DEFAULT TRUE, -- Кому видна поездка (всем или выбранным пользователям)
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания поездки
    user_id INTEGER NOT NULL, -- Идентификатор пользователя-создателя поездки
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
 );
