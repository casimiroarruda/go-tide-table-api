CREATE SCHEMA IF NOT EXISTS auth_store;

CREATE TABLE auth_store.clients (
    client_id UUID PRIMARY KEY,
    client_secret TEXT NOT NULL, -- Hash Bcrypt
    name TEXT NOT NULL,
    scopes TEXT[] NOT NULL,      -- Array de strings do Postgres
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);