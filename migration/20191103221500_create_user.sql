-- +migrate Up
CREATE TABLE IF NOT EXISTS "users" (
    "id" BIGINT NOT NULL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "role" INT NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS "users";
