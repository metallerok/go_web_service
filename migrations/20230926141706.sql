-- Create "users" table
CREATE TABLE "users" (
  "id" bigserial NOT NULL,
  "created_at" timestamp NULL,
  "updated_at" timestamp NULL,
  "deleted_at" timestamp NULL,
  "name" text NULL,
  "password" text NULL,
  "type" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
