DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_roles
        WHERE rolname = 'service'
    ) THEN
        CREATE USER service WITH PASSWORD 'test';
    END IF;
END $$;

GRANT CONNECT ON DATABASE trip TO service;
GRANT USAGE ON SCHEMA public TO service;

GRANT SELECT ON ALL TABLES IN SCHEMA public TO service;

GRANT INSERT, UPDATE, DELETE ON TABLE user, city, place, review, trip, trip_place, trip_photo, user_survey TO service;

ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO service;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT INSERT, UPDATE, DELETE ON TABLES TO service;

REVOKE ALL ON SCHEMA public FROM service;
REVOKE CREATE ON SCHEMA public FROM service;

-- Комментарий для лога (можно удалить или оставить для отладки)
DO $$ BEGIN
    RAISE NOTICE 'Service user created and privileges granted.';
END $$;
