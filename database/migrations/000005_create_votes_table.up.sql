CREATE TABLE votes(
  id SERIAL PRIMARY KEY,
  vote INTEGER NOT NULL CHECK(vote IN (1, -1)),
  user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE
  CASCADE,
  comment_id INTEGER NOT NULL REFERENCES comments(id) ON DELETE
  CASCADE,
  UNIQUE (user_id, comment_id)
);