CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  user_name VARCHAR UNIQUE NOT NULL,
  hashed_password CHAR(60) NOT NULL,
  full_name VARCHAR NOT NULL,
  role VARCHAR(30) NOT NULL DEFAULT 'Member',
  email VARCHAR UNIQUE NOT NULL,
  created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW()
);

ALTER TABLE users ADD CONSTRAINT unique_username_email UNIQUE (user_name, email);
