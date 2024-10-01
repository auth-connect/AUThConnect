CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "role" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "majors" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "university_name" varchar UNIQUE NOT NULL
);

CREATE TABLE "courses" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "semester" integer NOT NULL,
  "major_id" bigint NOT NULL
);

CREATE TABLE "forums" (
  "id" bigserial PRIMARY KEY,
  "course_id" bigserial NOT NULL
);

CREATE TABLE "threads" (
  "id" bigserial PRIMARY KEY,
  "forum_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "created_by" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "thread_id" bigint NOT NULL,
  "body" varchar,
  "created_by" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "archives" (
  "id" bigserial PRIMARY KEY,
  "course_id" bigserial NOT NULL
);

CREATE TABLE "files" (
  "id" bigserial PRIMARY KEY,
  "archive_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "created_by" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "courses" ADD FOREIGN KEY ("major_id") REFERENCES "majors" ("id");

ALTER TABLE "forums" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "threads" ADD FOREIGN KEY ("forum_id") REFERENCES "forums" ("id");

ALTER TABLE "threads" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("thread_id") REFERENCES "threads" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id");

ALTER TABLE "archives" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "files" ADD FOREIGN KEY ("archive_id") REFERENCES "archives" ("id");

ALTER TABLE "files" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id");
