CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL CHECK (LENGTH(first_name) >= 2),
    last_name VARCHAR(100) NOT NULL CHECK (LENGTH(last_name) >= 2),
    email VARCHAR(255) NOT NULL UNIQUE CHECK (POSITION('@' IN email) > 1),
    password VARCHAR(255) NOT NULL CHECK (LENGTH(password) >= 8),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    role VARCHAR(5) NOT NULL CHECK (role IN ('admin', 'user')) DEFAULT 'user',
    status VARCHAR(8) CHECK (status IN ('active', 'inactive')) DEFAULT 'inactive'
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_trigger WHERE tgname = 'update_user_updated_at') THEN
        CREATE TRIGGER update_user_updated_at
        BEFORE UPDATE ON users
        FOR EACH ROW
        EXECUTE FUNCTION update_updated_at_column();
    END IF;
END;
$$;
