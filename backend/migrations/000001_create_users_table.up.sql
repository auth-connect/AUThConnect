CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY,
  username varchar UNIQUE NOT NULL,
  hashed_password varchar NOT NULL,
  full_name varchar NOT NULL,
  role varchar NOT NULL,
  email varchar UNIQUE NOT NULL,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

ALTER TABLE users ADD CONSTRAINT unique_username_email UNIQUE (username, email);
