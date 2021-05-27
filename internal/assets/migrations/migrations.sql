-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE blobs (id serial PRIMARY KEY, blob jsonb);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE blobs;

