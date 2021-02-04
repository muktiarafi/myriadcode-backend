CREATE VIEW top_users AS
WITH scores AS (
  SELECT user_id, COUNT(*) * 10 AS score
  FROM user_solved_problems
  GROUP BY user_id
)
SELECT users.id, nickname, image_path, score
FROM users
JOIN scores ON users.id = scores.user_id
ORDER BY score DESC
LIMIT 10;