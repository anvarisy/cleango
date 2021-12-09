CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "name" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "authentications" (
  "token" bigserial PRIMARY KEY, 
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
