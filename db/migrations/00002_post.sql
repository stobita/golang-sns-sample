-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE post (
  id int NOT NULL AUTO_INCREMENT,
  title varchar(255) NOT NULL,
  content text NOT NULL,
  user_id int NOT NULL,
  created_at datetime,
  updated_at datetime,
  PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE post;
