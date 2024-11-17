CREATE TABLE IF NOT EXISTS"users" (
  "id" integer UNIQUE PRIMARY KEY,
  "username" varchar UNIQUE  NOT NULL,
    "hashed_password" varchar NOT NULL,
  "full_name" varchar(255) NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "role" varchar NOT NULL,
  "created_at"timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "cars" (
  "id" integer UNIQUE PRIMARY KEY,
  "title" varchar NOT NULL,
  "model_id" varchar NOT NULL,
  "model_name" varchar NOT NULL,
  "owner_id" integer  NOT NULL,
  "color" varchar NOT NULL,
  "description" varchar NOT NULL,
  "tags" varchar[] NOT NULL,
  "images" varchar[],
  "metadata" varchar[],
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "cars" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");
