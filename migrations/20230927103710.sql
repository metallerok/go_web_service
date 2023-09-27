-- Create "users" table
CREATE TABLE "users" (
  "id" bigserial NOT NULL,
  "name" text NULL,
  "password" text NULL,
  "type" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
