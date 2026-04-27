CREATE SCHEMA IF NOT EXISTS auth_store;

CREATE TABLE auth_store.clients (
    client_id UUID PRIMARY KEY,
    client_secret TEXT NOT NULL, -- Hash Bcrypt
    name TEXT NOT NULL,
    scopes TEXT[] NOT NULL,      -- Array de strings do Postgres
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);