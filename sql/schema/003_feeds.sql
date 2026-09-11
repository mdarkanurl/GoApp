-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    name        TEXT        NOT NULL,
    create_at   TIMESTAMP   NOT NULL,
    update_at   TIMESTAMP   NOT NULL,
    URL TEXT UNIQUE NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
