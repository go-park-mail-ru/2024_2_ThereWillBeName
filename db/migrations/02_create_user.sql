
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_roles 
        WHERE rolname = 'attractions_service'
    ) THEN
        CREATE USER attractions_service WITH PASSWORD 'test';
    END IF;

    IF NOT EXISTS (
        SELECT 1 
        FROM pg_roles 
        WHERE rolname = 'user_service'
    ) THEN
        CREATE USER user_service WITH PASSWORD 'test';
    END IF;

    IF NOT EXISTS (
        SELECT 1
        FROM pg_roles 
        WHERE rolname = 'trip_service'
    ) THEN
        CREATE USER trip_service WITH PASSWORD 'test';
    END IF;
END $$;

GRANT CONNECT ON DATABASE trip TO attractions_service;
GRANT USAGE ON SCHEMA public TO attractions_service;

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE place, review, city TO attractions_service;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO attractions_service;


GRANT CONNECT ON DATABASE trip TO user_service;
GRANT USAGE ON SCHEMA public TO user_service;

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE user TO user_service;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO user_service;


GRANT CONNECT ON DATABASE trip TO trip_service;
GRANT USAGE ON SCHEMA public TO trip_service;

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE trip, trip_place, trip_photo TO trip_service;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO trip_service;

REVOKE CREATE ON SCHEMA public FROM attractions_service;
REVOKE CREATE ON SCHEMA public FROM user_service;
REVOKE CREATE ON SCHEMA public FROM trip_service;

-- Комментарий для лога (можно удалить или оставить для отладки)
DO $$ BEGIN
    RAISE NOTICE 'Service user created and privileges granted.';
END $$;
