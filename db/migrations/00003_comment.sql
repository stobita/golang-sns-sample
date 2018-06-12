-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE comment (
  id int NOT NULL AUTO_INCREMENT,
  content text NOT NULL,
  user_id int NOT NULL,
  post_id int NOT NULL,
  created_at datetime,
  updated_at datetime,
  PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE comment;
