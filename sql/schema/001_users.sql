-- +goose up

CREATE TABLE users (
    id          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    name        TEXT        NOT NULL,
    create_at   TIMESTAMP   NOT NULL,
    update_at   TIMESTAMP   NOT NULL
);

-- +goose down
DROP TABLE users;
