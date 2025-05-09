-- Schema for library application
-- Creating books table
CREATE TABLE "books" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar NOT NULL,
  "author" varchar NOT NULL,
  "isbn" varchar UNIQUE NOT NULL,
  "published_year" int NOT NULL,
  "price" decimal(10, 2) NOT NULL,
  "quantity" int NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

-- Creating users table
CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

-- Creating loans table for book borrows
CREATE TABLE "loans" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "book_id" int NOT NULL,
  "borrowed_at" timestamptz NOT NULL DEFAULT (now()),
  "due_date" timestamptz NOT NULL,
  "returned_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  
  FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE,
  FOREIGN KEY ("book_id") REFERENCES "books" ("id") ON DELETE CASCADE
);

-- Indexes
CREATE INDEX ON "books" ("title");
CREATE INDEX ON "users" ("email");
CREATE INDEX ON "loans" ("user_id");
CREATE INDEX ON "loans" ("book_id");
CREATE INDEX ON "loans" ("borrowed_at");
CREATE INDEX ON "loans" ("due_date");
CREATE INDEX ON "loans" ("returned_at" NULLS FIRST);