-- +migrate Up
INSERT INTO "users" VALUES (
    0,
    'admin admin',
    'admin@admin.com',
    '$2a$14$/gshFrcGL9T9ot5mXGvfdusmXZR0bMBvXmtRh7t7ROZrDzO0e6pcO',
    0
);

INSERT INTO "users" VALUES (
    1,
    'teacher admin',
    'teacher@admin.com',
    '$2a$14$/gshFrcGL9T9ot5mXGvfdusmXZR0bMBvXmtRh7t7ROZrDzO0e6pcO',
    1
);

INSERT INTO "users" VALUES (
    2,
    'student admin',
    'student@admin.com',
    '$2a$14$/gshFrcGL9T9ot5mXGvfdusmXZR0bMBvXmtRh7t7ROZrDzO0e6pcO',
    2
);

-- +migrate Down
DELETE FROM "public"."users";
