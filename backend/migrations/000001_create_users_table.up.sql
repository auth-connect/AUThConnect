CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR UNIQUE NOT NULL,
  hashed_password BYTEA NOT NULL,
  role VARCHAR(30) NOT NULL DEFAULT 'Member',
  email CITEXT UNIQUE NOT NULL,
  activated BOOL NOT NULL,
  version INTEGER NOT NULL DEFAULT 1,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

ALTER TABLE users ADD CONSTRAINT unique_name_email UNIQUE (name, email);
