CREATE TABLE "users" (
  "id" varchar(12) PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "place_of_birth" varchar NOT NULL,
  "date_of_birth" timestamptz NOT NULL DEFAULT (now()),
  "avatar" varchar NOT NULL
);

CREATE TABLE "authentications" (
  "token" bigserial PRIMARY KEY, 
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
