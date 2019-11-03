-- +migrate Up
INSERT INTO "users" VALUES (
    0,
    'admin admin',
    'admin@admin.com',
    '$2a$14$/gshFrcGL9T9ot5mXGvfdusmXZR0bMBvXmtRh7t7ROZrDzO0e6pcO',
    0
);

-- +migrate Down
DELETE "users" WHERE id = 0;
