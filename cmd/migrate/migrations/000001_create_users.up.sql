CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    name varchar(255) NOT NULL,
    email citext NOT NULL UNIQUE,
    password_hash bytea NOT NULL,
    username varchar(255) NOT NULL UNIQUE,
    created_at timestamptz NOT NULL DEFAULT NOW()

);
