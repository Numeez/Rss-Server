-- +goose Up
CREATE TABLE feedfollowers(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    feed_id UUID NOT NULL REFERENCES feeds(id)
);
-- +goose Down
DROP TABLE feedfollowers;